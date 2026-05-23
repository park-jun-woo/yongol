//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what writeSSaCContextForTable — 테이블 관련 SSaC 컨텍스트 기록

package agent

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func writeSSaCContextForTable(b *strings.Builder, specsDir string, lookup map[string]features.Feature, table string) {
	for op, feat := range lookup {
		if feat.Table == table {
			writeSSaCForOp(b, specsDir, op)
		}
	}
}
