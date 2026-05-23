//ff:func feature=agent type=helper control=sequence
//ff:what insertHurlBlock — 새 Hurl 요청 블록을 파일 끝에 추가

package agent

import (
	"fmt"
	"strings"
)

// insertHurlBlock appends a new request block to the end of the hurl file.
func insertHurlBlock(originalContent, newBlock string) (string, error) {
	if !containsHTTPMethodLine(newBlock) {
		return "", fmt.Errorf("new Hurl block is missing HTTP method line (GET/POST/PUT/DELETE/PATCH)")
	}
	content := strings.TrimRight(originalContent, "\n")
	newBlock = strings.TrimRight(newBlock, "\n")
	return content + "\n\n" + newBlock + "\n", nil
}
