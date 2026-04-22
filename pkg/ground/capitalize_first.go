//ff:func feature=rule type=util control=sequence
//ff:what capitalizeFirst — 문자열 첫 글자를 PascalCase 로 변환 ("hashPassword" → "HashPassword")
package ground

// capitalizeFirst converts "hashPassword" → "HashPassword".
func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-('a'-'A')) + s[1:]
	}
	return s
}
