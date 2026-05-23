//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what writeFeatureTableContext — 테이블 관련 feature 컨텍스트 기록

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func writeFeatureTableContext(b *strings.Builder, ff *features.FeaturesFile, table string) {
	if ff == nil || table == "" {
		return
	}
	var related []features.Feature
	for _, f := range ff.Features {
		if f.Table == table {
			related = append(related, f)
		}
	}
	if len(related) > 0 {
		b.WriteString("Related features:\n")
		for _, f := range related {
			fmt.Fprintf(b, "  - %s %s: %s\n", f.Op, f.Path, f.Desc)
		}
		b.WriteByte('\n')
	}
}
