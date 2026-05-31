//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what writeOneFeatureAndArtifacts — 단일 feature 의 writeOneFeature + writeServiceArtifacts 호출/검증 헬퍼
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeOneFeatureAndArtifacts invokes writeOneFeature and (when plans exist)
// writeServiceArtifacts for the first plan, failing the test on error.
func writeOneFeatureAndArtifacts(t *testing.T, feature string, plans []*ir.ServicePlan, reg ir.TypeRegistry) {
	t.Helper()
	if err := writeOneFeature(feature, plans, t.TempDir(), reg); err != nil {
		t.Fatalf("writeOneFeature(%s): %v", feature, err)
	}
	if len(plans) > 0 {
		if err := writeServiceArtifacts(plans[0], t.TempDir(), reg); err != nil {
			t.Fatalf("writeServiceArtifacts(%s): %v", plans[0].OperationID, err)
		}
	}
}
