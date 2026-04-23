//ff:func feature=migration type=util control=sequence
//ff:what splitState.finish — 마지막 statement (trailing `;` 없음) 처리 후 out 반환
package migration

import "strings"

// finish flushes any trailing statement and returns the accumulated slice.
func (st *splitState) finish() []string {
	if strings.TrimSpace(st.sb.String()) != "" {
		st.out = append(st.out, st.sb.String())
	}
	return st.out
}
