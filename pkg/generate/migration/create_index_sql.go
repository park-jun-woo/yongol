//ff:func feature=migration type=accessor control=sequence
//ff:what CreateIndex.SQL — CREATE [UNIQUE] INDEX ... ON ... (...) [WHERE ...]
package migration

import (
	"fmt"
	"strings"
)

func (op CreateIndex) SQL() string {
	uniq := ""
	if op.Index.Unique {
		uniq = "UNIQUE "
	}
	using := ""
	if op.Index.Method != "" {
		using = " USING " + op.Index.Method
	}
	where := ""
	if op.Index.Where != "" {
		where = " WHERE " + op.Index.Where
	}
	return fmt.Sprintf("CREATE %sINDEX %s ON %s%s (%s)%s;",
		uniq, op.Index.Name, op.Table, using,
		strings.Join(op.Index.Columns, ", "), where)
}
