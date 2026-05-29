//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildMiddlewarePlanMinimal -- 최소 manifest → MiddlewarePlan 기본값 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildMiddlewarePlanMinimal(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Module: "github.com/test/proj"},
		},
	}
	ps := &prepared.State{}

	plan := BuildMiddlewarePlan(fs, ps)

	if !plan.RequestID {
		t.Error("RequestID should always be true")
	}
	if !plan.Prometheus {
		t.Error("Prometheus should be true by default")
	}
	if plan.BodyLimit == nil {
		t.Error("BodyLimit should never be nil")
	}
	if plan.ErrorEnvelope == nil {
		t.Error("ErrorEnvelope should never be nil")
	}
	if plan.RequestValidator == nil {
		t.Error("RequestValidator should never be nil")
	}
	if plan.SecurityHeaders == nil {
		t.Error("SecurityHeaders should not be nil by default")
	}
	if plan.CSRF != nil {
		t.Error("CSRF should be nil for no-auth projects")
	}
	if plan.BearerAuth != nil {
		t.Error("BearerAuth should be nil for no-auth projects")
	}
	if plan.RateLimit != nil {
		t.Error("RateLimit should be nil when no rules defined")
	}
}
