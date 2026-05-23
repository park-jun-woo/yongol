//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what writeOpenAPIPathContext — 테이블 관련 OpenAPI path 컨텍스트 기록

package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func writeOpenAPIPathContext(b *strings.Builder, specsDir string, lookup map[string]features.Feature, table string) {
	openapiPath := filepath.Join(specsDir, "api", "openapi.yaml")
	data, err := os.ReadFile(openapiPath)
	if err != nil {
		return
	}
	content := string(data)
	for op, feat := range lookup {
		if feat.Table == table {
			block := extractPathBlockForOp(content, op)
			if block != "" {
				fmt.Fprintf(b, "OpenAPI path block (%s):\n%s\n", op, block)
			}
		}
	}
}
