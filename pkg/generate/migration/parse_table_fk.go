//ff:func feature=migration type=parser control=sequence
//ff:what parseTableFK — 테이블 레벨 FOREIGN KEY (...) REFERENCES ... 파싱
package migration

import "strings"

// parseTableFK parses a table-scoped FOREIGN KEY clause.
// Returns nil when the clause is malformed.
func parseTableFK(t *Table, item string) *ForeignKey {
	rest := afterKeyword(item, "FOREIGN KEY")
	toks := tokenizeColumnDef(rest)
	if len(toks) < 3 {
		return nil
	}
	localCols := parseColumnList(toks[0])
	if strings.ToUpper(toks[1]) != "REFERENCES" {
		return nil
	}
	refTable, refCols, consumed := parseTableFKRef(toks)
	fk := &ForeignKey{
		Name:       FKName(t.Name, localCols),
		Columns:    localCols,
		RefTable:   refTable,
		RefColumns: refCols,
	}
	_ = applyRefOnActions(fk, toks, consumed)
	return fk
}
