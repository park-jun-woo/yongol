//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what isPureNumeric — 문자열이 순수 숫자/소수점/분수 형태인지 확인
package stml_design

// isPureNumeric returns true for pure numeric strings (integer, decimal, or fraction).
// E.g. "4", "0.5", "1/2"
func isPureNumeric(s string) bool {
	for _, c := range s {
		if (c >= '0' && c <= '9') || c == '.' || c == '/' {
			continue
		}
		return false
	}
	return true
}
