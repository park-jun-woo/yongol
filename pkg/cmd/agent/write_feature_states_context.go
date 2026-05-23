//ff:func feature=agent type=helper control=sequence
//ff:what writeFeatureStatesContext — 테이블의 states 컨텍스트 기록

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func writeFeatureStatesContext(b *strings.Builder, ff *features.FeaturesFile, table string) {
	if ff == nil || table == "" {
		return
	}
	td, ok := ff.Tables[table]
	if !ok || len(td.States) == 0 {
		return
	}
	fmt.Fprintf(b, "States for %s: %s\n\n", table, strings.Join(td.States, ", "))
}
