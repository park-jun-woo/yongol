//ff:func feature=gen-gogin type=util control=sequence
//ff:what truncateStderr — stderr 문자열을 limit 바이트 이내로 잘라낸다

package gogin

import "strings"

// truncateStderr trims stderr output to at most limit bytes, appending a
// marker when truncation occurs. Trailing whitespace is also stripped so
// error messages stay on a single line where possible.
func truncateStderr(s string, limit int) string {
	s = strings.TrimRight(s, "\n\r\t ")
	if limit <= 0 || len(s) <= limit {
		return s
	}
	return s[:limit] + "...(truncated)"
}
