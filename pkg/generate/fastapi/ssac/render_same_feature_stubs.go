//ff:func feature=gen-fastapi type=generator control=sequence
//ff:what RenderSameFeatureStubs — 같은 feature 내 @call/@eval 대상 inline stub 블록 생성

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// RenderSameFeatureStubs collects same-feature @call/@eval targets that lack
// definitions and returns inline stub functions to append at the bottom of
// the feature service file. Returns empty string when no stubs are needed.
func RenderSameFeatureStubs(plans []*ir.ServicePlan, feature string) string {
	stubNames := collectSameFeatureStubs(plans, feature)
	return renderInlineStubs(feature, stubNames)
}
