//ff:func feature=migration type=util control=iteration dimension=1
//ff:what stripLineComments — `--` 라인 코멘트 제거 (single-quoted 문자열 내부는 보존)
package migration

import "strings"

// stripLineComments removes SQL line comments introduced by `--`.
func stripLineComments(sql string) string {
	var sb strings.Builder
	lines := strings.Split(sql, "\n")
	for _, ln := range lines {
		idx := findLineCommentStart(ln)
		if idx >= 0 {
			sb.WriteString(ln[:idx])
		} else {
			sb.WriteString(ln)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
