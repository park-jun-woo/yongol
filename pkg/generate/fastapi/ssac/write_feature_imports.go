//ff:func feature=gen-fastapi type=generator control=sequence
//ff:what WriteFeatureImports — feature 단위 통합 import 블록 생성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// WriteFeatureImports writes the consolidated import block for a feature's
// service file, deduplicating across all plans. currentFeature is used for
// self-import prevention.
func WriteFeatureImports(plans []*ir.ServicePlan, currentFeature string) string {
	var b strings.Builder
	writeServiceImports(&b, plans, currentFeature)
	return b.String()
}
