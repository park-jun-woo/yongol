//ff:func feature=agent type=helper control=sequence
//ff:what stripMarkdownFences — LLM 출력에서 markdown 코드 펜스 제거

package agent

import "strings"

// stripMarkdownFences removes wrapping markdown code fences from LLM output.
func stripMarkdownFences(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "```") {
		return s
	}
	// Remove opening fence (``` or ```yaml etc.)
	idx := strings.Index(s, "\n")
	if idx < 0 {
		return s
	}
	s = s[idx+1:]

	// Remove closing fence
	if last := strings.LastIndex(s, "```"); last >= 0 {
		s = s[:last]
	}
	return strings.TrimSpace(s)
}
