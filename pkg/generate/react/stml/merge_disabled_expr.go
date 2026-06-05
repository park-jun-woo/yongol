//ff:func feature=stml-gen type=util control=sequence
//ff:what 기존 disabled 식(pending)에 data-enabled-when을 || !(<expr>)로 병합한다
package stml

// mergeDisabledExpr returns the inner JSX expression for a Button `disabled`
// attribute. base is the existing pending expression (e.g. "<mut>.isPending").
// When enabledWhen is set, the guard is OR-merged as `base || !(<expr>)` so the
// button is disabled while the action is not enabled. An empty enabledWhen
// returns base unchanged, preserving existing output.
func mergeDisabledExpr(base, enabledWhen, dataVar string) string {
	expr := renderEnabledWhenExpr(enabledWhen, dataVar)
	if expr == "" {
		return base
	}
	return base + " || !(" + expr + ")"
}
