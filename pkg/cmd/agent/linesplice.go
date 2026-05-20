//ff:func feature=agent type=helper control=pure
//ff:what spliceLines — 라인 범위 교체 공통 유틸

package agent

import "strings"

// spliceLines replaces lines[startLine:endLine] in content with replacement.
// Line numbers are 0-indexed. Returns the modified content.
func spliceLines(content string, startLine, endLine int, replacement string) string {
	lines := strings.Split(content, "\n")

	if startLine < 0 {
		startLine = 0
	}
	if endLine > len(lines) {
		endLine = len(lines)
	}
	if startLine > endLine {
		return content
	}

	// Build result: lines before + replacement + lines after
	var b strings.Builder
	for i := 0; i < startLine; i++ {
		b.WriteString(lines[i])
		b.WriteByte('\n')
	}

	rep := strings.TrimRight(replacement, "\n")
	if rep != "" {
		b.WriteString(rep)
		b.WriteByte('\n')
	}

	for i := endLine; i < len(lines); i++ {
		b.WriteString(lines[i])
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}

	return b.String()
}
