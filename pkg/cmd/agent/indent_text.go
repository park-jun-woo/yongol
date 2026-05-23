//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what indentText — 비어있지 않은 각 줄에 prefix 추가

package agent

import "strings"

func indentText(text, prefix string) string {
	lines := strings.Split(strings.TrimRight(text, "\n"), "\n")
	var b strings.Builder
	for _, line := range lines {
		if line == "" {
			b.WriteByte('\n')
		} else {
			b.WriteString(prefix)
			b.WriteString(line)
			b.WriteByte('\n')
		}
	}
	return b.String()
}
