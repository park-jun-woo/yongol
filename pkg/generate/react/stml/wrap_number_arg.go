//ff:func feature=stml-gen type=util control=sequence
//ff:what integer path param 표현식을 Number()로 래핑한다 — optional이면 null 가드 포함 (BUG-136)
package stml

// wrapNumberArg wraps an integer path-param expression in Number(). A required
// param is always present in the matched route, so it stays Number(expr). An
// optional param (":Name?") can be absent — Number(undefined)===NaN would be
// sent — so it becomes "expr != null ? Number(expr) : undefined".
func wrapNumberArg(expr string, optional bool) string {
	if optional {
		return expr + " != null ? Number(" + expr + ") : undefined"
	}
	return "Number(" + expr + ")"
}
