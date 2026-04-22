//ff:func feature=manifest type=util control=sequence
//ff:what applyVarcharLen — VARCHAR(N) 길이를 Table에 설정
package ddl

func applyVarcharLen(t *Table, colName, colType string) {
	n := extractVarcharLen(colType)
	if n <= 0 {
		return
	}
	if t.VarcharLen == nil {
		t.VarcharLen = map[string]int{}
	}
	t.VarcharLen[colName] = n
}
