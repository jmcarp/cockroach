// Code generated by "stringer -type=Key"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_1-0]
	_ = x[Start21_1PLUS-1]
	_ = x[Start21_2-2]
	_ = x[JoinTokensTable-3]
	_ = x[AcquisitionTypeInLeaseHistory-4]
	_ = x[SerializeViewUDTs-5]
	_ = x[ExpressionIndexes-6]
	_ = x[DeleteDeprecatedNamespaceTableDescriptorMigration-7]
	_ = x[FixDescriptors-8]
	_ = x[DatabaseRoleSettings-9]
	_ = x[TenantUsageTable-10]
	_ = x[SQLInstancesTable-11]
	_ = x[NewRetryableRangefeedErrors-12]
	_ = x[AlterSystemWebSessionsCreateIndexes-13]
	_ = x[SeparatedIntentsMigration-14]
	_ = x[PostSeparatedIntentsMigration-15]
	_ = x[RetryJobsWithExponentialBackoff-16]
	_ = x[RecordsBasedRegistry-17]
	_ = x[AutoSpanConfigReconciliationJob-18]
	_ = x[PreventNewInterleavedTables-19]
	_ = x[EnsureNoInterleavedTables-20]
	_ = x[DefaultPrivileges-21]
	_ = x[ZonesTableForSecondaryTenants-22]
	_ = x[UseKeyEncodeForHashShardedIndexes-23]
	_ = x[DatabasePlacementPolicy-24]
	_ = x[GeneratedAsIdentity-25]
	_ = x[OnUpdateExpressions-26]
	_ = x[SpanConfigurationsTable-27]
	_ = x[BoundedStaleness-28]
	_ = x[DateAndIntervalStyle-29]
	_ = x[PebbleFormatVersioned-30]
	_ = x[MarkerDataKeysRegistry-31]
	_ = x[PebbleSetWithDelete-32]
	_ = x[TenantUsageSingleConsumptionColumn-33]
	_ = x[SQLStatsTables-34]
	_ = x[SQLStatsCompactionScheduledJob-35]
	_ = x[V21_2-36]
	_ = x[Start22_1-37]
	_ = x[TargetBytesAvoidExcess-38]
}

const _Key_name = "V21_1Start21_1PLUSStart21_2JoinTokensTableAcquisitionTypeInLeaseHistorySerializeViewUDTsExpressionIndexesDeleteDeprecatedNamespaceTableDescriptorMigrationFixDescriptorsDatabaseRoleSettingsTenantUsageTableSQLInstancesTableNewRetryableRangefeedErrorsAlterSystemWebSessionsCreateIndexesSeparatedIntentsMigrationPostSeparatedIntentsMigrationRetryJobsWithExponentialBackoffRecordsBasedRegistryAutoSpanConfigReconciliationJobPreventNewInterleavedTablesEnsureNoInterleavedTablesDefaultPrivilegesZonesTableForSecondaryTenantsUseKeyEncodeForHashShardedIndexesDatabasePlacementPolicyGeneratedAsIdentityOnUpdateExpressionsSpanConfigurationsTableBoundedStalenessDateAndIntervalStylePebbleFormatVersionedMarkerDataKeysRegistryPebbleSetWithDeleteTenantUsageSingleConsumptionColumnSQLStatsTablesSQLStatsCompactionScheduledJobV21_2Start22_1TargetBytesAvoidExcess"

var _Key_index = [...]uint16{0, 5, 18, 27, 42, 71, 88, 105, 154, 168, 188, 204, 221, 248, 283, 308, 337, 368, 388, 419, 446, 471, 488, 517, 550, 573, 592, 611, 634, 650, 670, 691, 713, 732, 766, 780, 810, 815, 824, 846}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
