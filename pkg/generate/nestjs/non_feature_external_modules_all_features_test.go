//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNonFeatureExternalModulesAllFeatures — TestNonFeatureExternalModules — feature module 이 아닌 external package 이름만 반환 검증

package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestNonFeatureExternalModulesAllFeatures(t *testing.T) {
	extPkgs := []externalPackage{{Name: "course"}}
	plansByFeature := map[string][]*ir.ServicePlan{"course": {{}}}
	if got := nonFeatureExternalModules(extPkgs, plansByFeature); got != nil {
		t.Errorf("got %v, want nil when all external pkgs are features", got)
	}
}
