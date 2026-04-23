//ff:func feature=migration type=util control=selection
//ff:what mnemonicFor — 단일 Operation 에 해당하는 snake_case mnemonic 반환
package migration

import "fmt"

// mnemonicFor returns a short snake_case identifier describing op.
func mnemonicFor(op Operation) string {
	switch v := op.(type) {
	case CreateTable:
		return "create_" + sanitize(v.Table.Name)
	case DropTable:
		return "drop_" + sanitize(v.Name)
	case AddColumn:
		return "add_" + sanitize(v.Table) + "_" + sanitize(v.Column.Name)
	case DropColumn:
		return "drop_" + sanitize(v.Table) + "_" + sanitize(v.Column)
	case AlterColumnType:
		return "alter_" + sanitize(v.Table) + "_" + sanitize(v.Column) + "_type"
	case AlterColumnNullable:
		if v.To {
			return "nullable_" + sanitize(v.Table) + "_" + sanitize(v.Column)
		}
		return "notnull_" + sanitize(v.Table) + "_" + sanitize(v.Column)
	case AlterColumnDefault:
		return "default_" + sanitize(v.Table) + "_" + sanitize(v.Column)
	case CreateIndex:
		return "create_index_" + sanitize(v.Index.Name)
	case DropIndex:
		return "drop_index_" + sanitize(v.Name)
	case AddForeignKey:
		return "add_fk_" + sanitize(v.FK.Name)
	case DropForeignKey:
		return "drop_fk_" + sanitize(v.Name)
	case AddCheck:
		return "add_check_" + sanitize(v.Check.Name)
	case DropCheck:
		return "drop_check_" + sanitize(v.Name)
	case RenameColumn:
		return "rename_" + sanitize(v.Table) + "_" + sanitize(v.From) + "_to_" + sanitize(v.To)
	case RenameTable:
		return "rename_table_" + sanitize(v.From) + "_to_" + sanitize(v.To)
	}
	return fmt.Sprintf("change_%T", op)
}
