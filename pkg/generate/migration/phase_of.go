//ff:func feature=migration type=util control=selection
//ff:what phaseOf — Operation 타입을 1~12 단계 정수로 매핑
package migration

// phaseOf returns the sort phase (1..12) of an Operation.
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
	case AddColumn:
		return 8
	case AlterColumnType, AlterColumnNullable, AlterColumnDefault:
		return 9
	case AddCheck:
		return 10
	case CreateIndex:
		return 11
	case AddForeignKey:
		return 12
	}
	return 99
}
