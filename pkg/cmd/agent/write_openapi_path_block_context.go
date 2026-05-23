//ff:func feature=agent type=helper control=sequence
//ff:what writeOpenAPIPathBlockContext — opID의 OpenAPI path 블록 컨텍스트 기록

package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func writeOpenAPIPathBlockContext(b *strings.Builder, specsDir, opID string) {
	openapiPath := filepath.Join(specsDir, "api", "openapi.yaml")
	data, err := os.ReadFile(openapiPath)
	if err != nil {
		return
	}
	block := extractPathBlockForOp(string(data), opID)
	if block != "" {
		fmt.Fprintf(b, "OpenAPI path block (%s):\n%s\n", opID, block)
	}
}
