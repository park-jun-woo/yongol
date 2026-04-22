//ff:func feature=migration type=util control=sequence
//ff:what InferDescription — ops 리스트에서 파일명용 snake_case 요약 생성
package migration

import (
	"fmt"
	"strings"
)

// InferDescription returns a snake_case string used as the <desc> in
// artifacts/db/migrations/NNNN_<desc>.sql. Rules:
//   - empty ops        -> "noop"
//   - single op        -> mnemonic derived from op.Description()
//   - multiple ops     -> "<mnemonic-of-first>_and_more"
//   - initial (single CreateTable) handled by caller (InitialMigrationDesc)
func InferDescription(ops []Operation) string {
	if len(ops) == 0 {
		return "noop"
	}
	mnemonic := mnemonicFor(ops[0])
	if len(ops) == 1 {
		return mnemonic
	}
	// All creates → initial-ish bundle
	allCreates := true
	for _, op := range ops {
		if _, ok := op.(CreateTable); !ok {
			allCreates = false
			break
		}
	}
	if allCreates {
		return "initial"
	}
	return mnemonic + "_and_more"
}

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

// sanitize keeps [a-z0-9_] and replaces others with `_`.
func sanitize(s string) string {
	s = strings.ToLower(s)
	b := strings.Builder{}
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '_':
			b.WriteRune('_')
		default:
			b.WriteRune('_')
		}
	}
	return b.String()
}
