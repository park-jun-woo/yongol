//ff:func feature=migration type=util control=sequence
//ff:what trimForErr — 오류 메시지용 SQL 조각 축약 (개행 대체 + 80자 컷)
package migration

import "strings"

// trimForErr sanitises a SQL fragment for use inside an error message.
func trimForErr(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > 80 {
		return s[:80] + "..."
	}
	return s
}
