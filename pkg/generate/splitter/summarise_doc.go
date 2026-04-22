//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what summariseDoc — doc comment 에서 첫 번째 비어있지 않은 라인을 추출해 //ff:what 요약으로 사용
package splitter

import "strings"

// summariseDoc picks the first non-blank line from a doc comment and
// normalises it to a single-line //ff:what summary. Leading "//" prefixes
// and whitespace are stripped; when the doc is empty it falls back to
// the declaration identifier so the annotation stays non-empty (A3).
func summariseDoc(doc, fallback string) string {
	for _, line := range strings.Split(doc, "\n") {
		trimmed := strings.TrimSpace(line)
		trimmed = strings.TrimPrefix(trimmed, "//")
		trimmed = strings.TrimSpace(trimmed)
		if trimmed == "" {
			continue
		}
		return trimmed
	}
	if fallback != "" {
		return fallback
	}
	return "generated"
}
