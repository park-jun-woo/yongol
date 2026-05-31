//ff:func feature=gen-nestjs type=test-helper control=iteration dimension=1
//ff:what exerciseOneFeature — plansByFeature 첫 feature 에 대해 writeOneFeature/writeServiceArtifacts 직접 호출 헬퍼
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// exerciseOneFeature runs writeOneFeature and writeServiceArtifacts for the
// first feature so coverage attribution reaches those names.
func exerciseOneFeature(t *testing.T, plansByFeature map[string][]*ir.ServicePlan, reg ir.TypeRegistry) {
	for feature, plans := range plansByFeature {
		writeOneFeatureAndArtifacts(t, feature, plans, reg)
		return
	}
}
