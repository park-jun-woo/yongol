//ff:func feature=migration type=util control=selection
//ff:what phaseOf — Operation 타입을 1~12 단계 정수로 매핑
package migration

// phaseOf returns the sort phase (1..13) of an Operation.
func phaseOf(op Operation) int {
	switch op.(type) {
	case RenameTable, RenameColumn:
		return 1
	case DropForeignKey:
		return 2
	case DropIndex:
		return 3
	case DropCheck:
		return 4
	case DropColumn:
		return 5
	case DropTable:
		return 6
	case CreateTable:
		return 7
	case InsertSentinel:
		return 8
	case AddColumn:
		return 9
	case AlterColumnType, AlterColumnNullable, AlterColumnDefault:
		return 10
	case AddCheck:
		return 11
	case CreateIndex:
		return 12
	case AddForeignKey:
		return 13
	}
	return 99
}
