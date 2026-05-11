//ff:func feature=stml-gen type=util control=sequence
//ff:what 빈 문자열이면 기본값을 반환한다
package stml

func orDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}
