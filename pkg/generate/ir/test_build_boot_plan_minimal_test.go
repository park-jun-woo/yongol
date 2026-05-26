//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildBootPlanMinimal -- 최소 manifest → BootPlan 기본 블록 활성화 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildBootPlanMinimal(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Metadata: manifest.Metadata{Name: "testproj"},
			Backend:  manifest.Backend{Module: "github.com/test/proj"},
		},
	}
	ps := &prepared.State{}

	plan := BuildBootPlan(fs, ps, "go-gin")

	if plan.ProjectID != "testproj" {
		t.Errorf("ProjectID = %q, want %q", plan.ProjectID, "testproj")
	}
	if plan.Module != "github.com/test/proj" {
		t.Errorf("Module = %q, want %q", plan.Module, "github.com/test/proj")
	}
	if plan.BackendType != "go-gin" {
		t.Errorf("BackendType = %q, want %q", plan.BackendType, "go-gin")
	}
	if plan.ErrorEnvelope == nil {
		t.Fatal("ErrorEnvelope is nil")
	}
	if len(plan.ErrorEnvelope.Fields) != 3 {
		t.Errorf("ErrorEnvelope.Fields len = %d, want 3", len(plan.ErrorEnvelope.Fields))
	}
	if len(plan.ActiveBlocks) != 25 {
		t.Errorf("ActiveBlocks len = %d, want 25", len(plan.ActiveBlocks))
	}
}
