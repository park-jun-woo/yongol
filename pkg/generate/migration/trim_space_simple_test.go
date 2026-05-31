//ff:func feature=migration type=test control=iteration dimension=1
//ff:what tokenize_split_unit_test — tokenizeColumnDef/splitStatements/splitTopLevel/parseColumnList/collectDefaultExpr 단위 테스트
package migration

func trimSpaceSimple(s string) string {
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\n' || s[start] == '\t' || s[start] == '\r') {
		start++
	}
	end := len(s)
	for end > start && (s[end-1] == ' ' || s[end-1] == '\n' || s[end-1] == '\t' || s[end-1] == '\r') {
		end--
	}
	return s[start:end]
}
