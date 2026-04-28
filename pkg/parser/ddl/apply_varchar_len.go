//ff:func feature=manifest type=util control=sequence
//ff:what applyVarcharLen — VARCHAR(N) 길이를 Column에 설정
package ddl

func applyVarcharLen(col *Column, colType string) {
	n := extractVarcharLen(colType)
	if n <= 0 {
		return
	}
	col.VarcharLen = n
}
