//ff:func feature=stml-gen type=util control=selection dimension=1
//ff:what 가드 비교 연산자를 JS 연산자로 매핑한다 (= → ===, != → !==, 나머지 유지)
package stml

// jsxCompareOp maps a guard comparison operator to its JavaScript equivalent.
// "=" becomes strict equality "===" and "!=" becomes strict inequality "!==";
// relational operators (>, <, >=, <=) pass through unchanged.
func jsxCompareOp(op string) string {
	switch op {
	case "=":
		return "==="
	case "!=":
		return "!=="
	default:
		return op
	}
}
