//ff:func feature=migration type=parser control=sequence
//ff:what parseInlineRef — 컬럼 뒤 REFERENCES <table>(<col>) [ON DELETE ...] 파싱
package migration

// parseInlineRef parses the `REFERENCES ...` clause attached to a column.
// Returns the ForeignKey and number of tokens consumed from `toks`.
func parseInlineRef(t *Table, col string, toks []string) (*ForeignKey, int) {
	if len(toks) == 0 {
		return nil, 0
	}
	refTable, refCol, consumed := parseInlineRefTarget(toks)
	fk := &ForeignKey{
		Name:       FKName(t.Name, []string{col}),
		Columns:    []string{col},
		RefTable:   refTable,
		RefColumns: []string{refCol},
	}
	consumed = applyRefOnActions(fk, toks, consumed)
	return fk, consumed
}
