//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNonFeatureExternalModules — TestNonFeatureExternalModules — feature module 이 아닌 external package 이름만 반환 검증

package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestNonFeatureExternalModules(t *testing.T) {
	extPkgs := []externalPackage{
		{Name: "billing"}, // not a feature -> included
		{Name: "course"},  // also a feature -> excluded
		{Name: "mailer"},  // not a feature -> included
	}
	plansByFeature := map[string][]*ir.ServicePlan{
		"course": {{}},
	}

	got := nonFeatureExternalModules(extPkgs, plansByFeature)
	if len(got) != 2 {
		t.Fatalf("got %v, want 2 non-feature modules", got)
	}
	if got[0] != "billing" || got[1] != "mailer" {
		t.Errorf("got %v, want [billing mailer]", got)
	}
}
