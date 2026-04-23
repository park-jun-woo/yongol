//ff:func feature=migration type=util control=iteration dimension=1
//ff:what splitTopLevel — paren-depth 0 기준으로 문자열 분할 (따옴표 내부는 건너뜀)
package migration

// splitTopLevel splits s on `sep` at top paren-depth zero.
func splitTopLevel(s string, sep byte) []string {
	st := newSplitState()
	for i := 0; i < len(s); i++ {
		i = stepTopLevel(st, s, i, sep)
	}
	if st.sb.Len() > 0 {
		st.out = append(st.out, st.sb.String())
	}
	return st.out
}
