//ff:func feature=agent type=helper control=sequence
//ff:what writeSSaCForOp — operationId의 SSaC 파일 컨텍스트 기록

package agent

import (
	"fmt"
	"strings"
)

func writeSSaCForOp(b *strings.Builder, specsDir, opID string) {
	content, ok := findSSaCFile(specsDir, opID)
	if ok && content != "" {
		fmt.Fprintf(b, "SSaC (%s.ssac):\n%s\n\n", opID, content)
	}
}
