//ff:func feature=external type=util control=sequence
//ff:what readBodySnippet — response body 를 limit byte 까지 읽어 단일 줄 문자열로 반환

package external

import (
	"io"
	"strings"
)

// readBodySnippet — reads up to limit bytes from a response body and returns a single-line string (best-effort, errors ignored)
func readBodySnippet(r io.Reader, limit int) string {
	buf := make([]byte, limit+1)
	n, _ := io.ReadFull(r, buf)
	if n == 0 {
		return ""
	}
	truncated := n > limit
	if truncated {
		n = limit
	}
	s := strings.ReplaceAll(strings.TrimSpace(string(buf[:n])), "\n", " ")
	if truncated {
		s += "..."
	}
	return s
}
