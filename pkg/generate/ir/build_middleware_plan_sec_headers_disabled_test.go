//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildMiddlewarePlanSecHeadersDisabled -- security headers 명시 비활성화 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildMiddlewarePlanSecHeadersDisabled(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module:          "github.com/test/nosh",
				SecurityHeaders: &manifest.SecurityHeadersConfig{Enabled: &disabled},
			},
		},
	}
	ps := &prepared.State{}

	plan := BuildMiddlewarePlan(fs, ps)

	if plan.SecurityHeaders != nil {
		t.Error("SecurityHeaders should be nil when explicitly disabled")
	}
}
