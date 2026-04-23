//ff:func feature=migration type=parser control=sequence
//ff:what parseColumn — CREATE TABLE(...) 내부 단일 컬럼 정의 파싱
package migration

// parseColumn handles a single column definition item inside CREATE TABLE(...).
func parseColumn(t *Table, def string) {
	tokens := tokenizeColumnDef(def)
	if len(tokens) < 2 {
		return
	}
	name := canonIdent(tokens[0])
	typeTok, rest := collectTypeTokens(tokens[1:])
	ct, isSerial := NormalizeType(typeTok)

	col := &Column{Name: name, Type: ct, Nullable: true}
	applyColumnAttrs(t, col, rest)
	applySerialDefault(t, col, isSerial)

	t.Columns = append(t.Columns, col)
}
