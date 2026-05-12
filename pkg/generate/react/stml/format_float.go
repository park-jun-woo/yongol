//ff:func feature=stml-gen type=util control=sequence
//ff:what float64를 JS 출력용 문자열로 포맷한다 (정수값이면 소수점 제거)
package stml

import "fmt"

// formatFloat formats a float64 for JS output, dropping the decimal point for
// integer values (e.g. 100.0 → "100", 3.14 → "3.14").
func formatFloat(v float64) string {
	if v == float64(int64(v)) {
		return fmt.Sprintf("%d", int64(v))
	}
	return fmt.Sprintf("%g", v)
}
