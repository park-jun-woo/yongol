//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteByName_ZeroCov — buildPlansByFeature/writeFeatureModules/writeOneFeature/writeServiceArtifacts 직접 호출
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildPlansByFeature_Empty_ZeroCov(t *testing.T) {
	// No service funcs → empty map (loop body never runs).
	out := buildPlansByFeature(&yongol.Fullstack{})
	if len(out) != 0 {
		t.Errorf("empty fs should yield empty map, got %v", out)
	}
}
