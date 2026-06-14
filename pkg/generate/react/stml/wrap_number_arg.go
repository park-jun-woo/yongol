//ff:func feature=stml-gen type=util control=sequence
//ff:what integer path param 표현식을 평문 Number()로 래핑한다 (BUG-137)
package stml

// wrapNumberArg wraps an integer path-param expression in Number(). Number()'s
// return type is `number` even when expr is `string | undefined`, so a plain
// Number(expr) is always type-compatible with a required path parameter
// (BUG-137: a `: undefined` branch widened the type to `number | undefined`,
// causing TS2322). Runtime NaN from an absent optional segment is prevented by
// a separate call-guard (query `enabled` / mutation trigger disabled), not by
// mutilating the arg (BUG-136).
func wrapNumberArg(expr string) string {
	return "Number(" + expr + ")"
}
