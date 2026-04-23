//ff:func feature=migration type=util control=selection
//ff:what checkSafetyForOp — 단일 Operation 에 대한 SafetyIssue 리스트 생성 (type switch)
package migration

// checkSafetyForOp dispatches to rule-specific safety checks.
func checkSafetyForOp(op Operation) []SafetyIssue {
	switch v := op.(type) {
	case AlterColumnNullable:
		return safetyNotNullWithoutBackfill(v)
	case AddColumn:
		return safetyAddColumnNotNull(v)
	case DropTable:
		return safetyDropTable(v)
	case DropColumn:
		return safetyDropColumn(v)
	case AlterColumnType:
		return safetyAlterColumnType(v)
	}
	return nil
}
