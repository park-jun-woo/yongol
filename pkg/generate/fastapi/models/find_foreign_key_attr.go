//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what findForeignKeyAttr — ForeignKey 배열에서 컬럼 매칭 FK 속성 문자열 검색

package models

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// findForeignKeyAttr returns a ForeignKey("table.col") attribute string
// when the column has a foreign key reference, or "" otherwise.
func findForeignKeyAttr(fks []ddl.ForeignKey, colName string) string {
	for _, fk := range fks {
		if fk.Column == colName {
			return fmt.Sprintf("ForeignKey(\"%s.%s\")", fk.RefTable, fk.RefColumn)
		}
	}
	return ""
}
