// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package scexec

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/jobs/jobspb"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descpb"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/tabledesc"
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scop"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/errors"
)

// ExecuteOps executes the provided ops. The ops must all be of the same type.
func ExecuteOps(ctx context.Context, deps Dependencies, toExecute scop.Ops) error {
	log.Infof(ctx, "executing %d ops of type %s", len(toExecute.Slice()), toExecute.Type().String())

	if deps.TestingKnobs() != nil && deps.TestingKnobs().BeforeStage != nil {
		md := TestingKnobMetadata{
			Statements: deps.Statements(),
			Phase:      deps.Phase(),
		}
		if err := deps.TestingKnobs().BeforeStage(toExecute, md); err != nil {
			return err
		}
	}
	switch typ := toExecute.Type(); typ {
	case scop.MutationType:
		return executeDescriptorMutationOps(ctx, deps, toExecute.Slice())
	case scop.BackfillType:
		return executeBackfillOps(ctx, deps, toExecute.Slice())
	case scop.ValidationType:
		return executeValidationOps(ctx, deps, toExecute.Slice())
	default:
		return errors.AssertionFailedf("unknown ops type %d", typ)
	}
}

// UpdateDescriptorJobIDs updates the job ID for the schema change on the
// specified set of table descriptors.
func UpdateDescriptorJobIDs(
	ctx context.Context,
	catalog Catalog,
	descIDs []descpb.ID,
	expectedID jobspb.JobID,
	newID jobspb.JobID,
) error {
	b := catalog.NewCatalogChangeBatcher()
	for _, id := range descIDs {
		// Confirm the descriptor is a table, view or sequence
		// since we can only lock those types.
		desc, err := catalog.MustReadMutableDescriptor(ctx, id)
		if err != nil {
			return err
		}

		// Currently all "locking" schema changes are on tables. This will probably
		// need to be expanded at least to types.
		table, ok := desc.(*tabledesc.Mutable)
		if !ok {
			continue
		}
		if oldID := jobspb.JobID(table.NewSchemaChangeJobID); oldID != expectedID {
			return errors.AssertionFailedf(
				"unexpected schema change job ID %d on table %d, expected %d", oldID, table.GetID(), expectedID)
		}
		table.NewSchemaChangeJobID = int64(newID)
		if err := b.CreateOrUpdateDescriptor(ctx, table); err != nil {
			return err
		}
	}
	return b.ValidateAndRun(ctx)
}
