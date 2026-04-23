//ff:func feature=orchestrator type=parser control=sequence
//ff:what parseSQLCLine — 한 줄에서 sqlc name 주석을 파싱해 QuerySpec 반환

package sqlc

// parseSQLCLine extracts a QuerySpec from a single sqlc source line when it
// matches the `-- name: Foo :cardinality` pattern. Returns (spec, true) on
// match; the caller is responsible for file path and line number.
//
// RowType follows sqlc's default naming convention: "<Name>Row" for ":one" /
// ":many" queries. ":exec" / ":execresult" do not produce a row struct, so
// RowType is left empty.
func parseSQLCLine(line, model, file string, lineNo int) (QuerySpec, bool) {
	m := sqlcNameRe.FindStringSubmatch(line)
	if m == nil {
		return QuerySpec{}, false
	}
	name := m[1]
	cardinality := ""
	if len(m) >= 3 {
		cardinality = m[2]
	}
	return QuerySpec{
		Name:        name,
		Model:       model,
		Method:      stripModelPrefix(name, model),
		Cardinality: cardinality,
		RowType:     rowTypeFor(name, cardinality),
		File:        file,
		Line:        lineNo,
	}, true
}
