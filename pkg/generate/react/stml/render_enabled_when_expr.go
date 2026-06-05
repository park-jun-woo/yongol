//ff:func feature=stml-gen type=util control=sequence
//ff:what data-enabled-when 가드를 disabled 병합용 JSX boolean 표현식으로 변환한다
package stml

// renderEnabledWhenExpr converts a data-enabled-when guard condition into a JSX
// boolean expression by reusing the Phase001 guard converter
// (resolveStateCondition). The condition follows the same dataVar binding rule
// as data-state: the guard's model prefix is a top-level property of the fetched
// resource bound under dataVar. Returns "" for an empty condition so callers can
// keep their existing output unchanged.
func renderEnabledWhenExpr(enabledWhen, dataVar string) string {
	if enabledWhen == "" {
		return ""
	}
	return resolveStateCondition(enabledWhen, dataVar)
}
