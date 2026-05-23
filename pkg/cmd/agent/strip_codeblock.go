//ff:func feature=agent type=helper control=sequence
//ff:what stripCodeBlock — LLM 응답에서 markdown 코드블록 울타리 제거

package agent

import "strings"

// stripCodeBlock removes markdown code block fences from LLM output.
// Handles ```sql ... ```, ```yaml ... ```, ``` ... ``` etc.
// Returns the inner content with surrounding whitespace trimmed.
func stripCodeBlock(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "```") {
		return s
	}

	// Remove opening fence line (```sql, ```yaml, ``` etc.)
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
