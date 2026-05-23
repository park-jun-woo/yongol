//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what extractHurlBlock — Hurl 요청 블록 단위 추출 (# OperationId 기준)

package agent

import (
	"fmt"
	"strings"
)

// extractHurlBlock extracts the request block starting with "# OperationId" comment.
// Block ends at next "# " comment or EOF.
func extractHurlBlock(content, operationId string) (block string, startLine, endLine int, err error) {
	lines := strings.Split(content, "\n")

	commentPrefix := "# " + operationId
	blockStart := -1
	for i, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == commentPrefix {
			blockStart = i
			break
		}
	}
	if blockStart < 0 {
		return "", 0, 0, fmt.Errorf("operationId %q not found as Hurl comment (# %s)", operationId, operationId)
	}

	blockEnd := len(lines)
	for i := blockStart + 1; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmed, "# ") && !strings.HasPrefix(trimmed, "# @") {
			blockEnd = i
			break
		}
	}

	for blockEnd > blockStart+1 && strings.TrimSpace(lines[blockEnd-1]) == "" {
		blockEnd--
	}

	startLine = blockStart
	endLine = blockEnd

	var b strings.Builder
	for i := startLine; i < endLine; i++ {
		b.WriteString(lines[i])
		if i < endLine-1 {
			b.WriteByte('\n')
		}
	}

	return b.String(), startLine, endLine, nil
}
