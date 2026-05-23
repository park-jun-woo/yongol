//ff:func feature=agent type=helper control=sequence
//ff:what writeFeaturePublicContext — feature public 속성 컨텍스트 기록

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func writeFeaturePublicContext(b *strings.Builder, feat features.Feature) {
	fmt.Fprintf(b, "Feature public: %v\n\n", feat.Public)
}
