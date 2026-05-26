//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildMiddlewarePlanSecHeadersDev -- dev 프로파일 → CSPReportOnly 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildMiddlewarePlanSecHeadersDev(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module: "github.com/test/dev",
				SecurityHeaders: &manifest.SecurityHeadersConfig{
					Profile: "dev",
				},
			},
		},
	}
	ps := &prepared.State{}

	plan := BuildMiddlewarePlan(fs, ps)

	if plan.SecurityHeaders == nil {
		t.Fatal("SecurityHeaders should not be nil")
	}
	if plan.SecurityHeaders.Profile != "dev" {
		t.Errorf("SecurityHeaders.Profile = %q, want %q", plan.SecurityHeaders.Profile, "dev")
	}
	if !plan.SecurityHeaders.CSPReportOnly {
		t.Error("dev profile should set CSPReportOnly = true")
	}
}
