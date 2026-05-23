//ff:func feature=agent type=helper control=sequence
//ff:what insertRegoBlock — 새 Rego allow 블록을 파일 끝에 추가

package agent

import (
	"fmt"
	"strings"
)

// insertRegoBlock appends a new allow block to the end of the rego file.
func insertRegoBlock(originalContent, newBlock string) (string, error) {
	if !strings.Contains(newBlock, "allow") || !strings.Contains(newBlock, "if") || !strings.Contains(newBlock, "{") {
		return "", fmt.Errorf("new Rego block is missing 'allow if {' pattern")
	}
	content := strings.TrimRight(originalContent, "\n")
	newBlock = strings.TrimRight(newBlock, "\n")
	return content + "\n\n" + newBlock + "\n", nil
}
