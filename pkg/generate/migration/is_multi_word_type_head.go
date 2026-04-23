//ff:func feature=migration type=util control=selection
//ff:what isMultiWordTypeHead — CHARACTER/TIMESTAMP/TIME/DOUBLE 처럼 후속 토큰이 이어질 수 있는 타입 헤드 판정
package migration

// isMultiWordTypeHead reports whether upper is one of the SQL type
// prefixes that may pair with VARYING / PRECISION / WITH / WITHOUT.
func isMultiWordTypeHead(upper string) bool {
	switch upper {
	case "CHARACTER", "TIMESTAMP", "TIME", "DOUBLE":
		return true
	}
	return false
}
