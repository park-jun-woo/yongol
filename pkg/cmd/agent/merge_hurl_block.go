//ff:func feature=agent type=helper control=sequence
//ff:what mergeHurlBlock — Hurl 요청 블록을 원본 콘텐츠에 머지

package agent

import "fmt"

// mergeHurlBlock replaces the request block in original content.
func mergeHurlBlock(originalContent string, startLine, endLine int, fixedBlock string) (string, error) {
	if !containsHTTPMethodLine(fixedBlock) {
		return "", fmt.Errorf("fixed Hurl block is missing HTTP method line (GET/POST/PUT/DELETE/PATCH)")
	}
	return spliceLines(originalContent, startLine, endLine, fixedBlock), nil
}
