//ff:func feature=gen-ir type=util control=selection
//ff:what isCountResultType -- 스칼라 정수 result 타입(COUNT 쿼리) 여부 판정

package ir

// isCountResultType returns true for scalar integer result types that indicate
// a COUNT query rather than a row query.
func isCountResultType(t string) bool {
	switch t {
	case "int64", "int32", "int", "uint64":
		return true
	}
	return false
}
