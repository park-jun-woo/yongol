//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectModuleDepsEmpty — TestCollectModuleDeps — plan들에서 queue/authz/same-feature-stub + 정렬된 cross-feature @call 수집 검증

package ssac

import (
	"testing"
)

func TestCollectModuleDepsEmpty(t *testing.T) {
	deps := collectModuleDeps("Course", nil)
	if deps.NeedsQueue || deps.NeedsAuthz || deps.NeedsSameFeatureStub {
		t.Errorf("expected no deps, got %+v", deps)
	}
	if len(deps.CrossFeatures) != 0 {
		t.Errorf("CrossFeatures = %v, want empty", deps.CrossFeatures)
	}
}
