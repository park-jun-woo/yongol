//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteByName_ZeroCov — buildPlansByFeature/writeFeatureModules/writeOneFeature/writeServiceArtifacts 직접 호출
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/types"
)

func TestWriteFeatureModules_Empty_ZeroCov(t *testing.T) {
	// Empty plansByFeature → no features written, nil error.
	names, err := writeFeatureModules(map[string][]*ir.ServicePlan{}, t.TempDir(), types.NewRegistry())
	if err != nil {
		t.Fatalf("writeFeatureModules empty: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("expected no feature names, got %v", names)
	}
}
