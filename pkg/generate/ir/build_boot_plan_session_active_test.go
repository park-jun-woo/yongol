//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestBuildBootPlanSessionActive -- session 활성 / cache 비활성 블록 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildBootPlanSessionActive(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Metadata: manifest.Metadata{Name: "sessproj"},
			Backend:  manifest.Backend{Module: "github.com/test/sess"},
		},
	}
	ps := &prepared.State{
		ActiveBackends: prepared.ActiveBackends{
			Session: &prepared.Session{Backend: "postgres"},
		},
	}

	plan := BuildBootPlan(fs, ps, "go-gin")

	blockMap := map[string]bool{}
	for _, b := range plan.ActiveBlocks {
		blockMap[b.Name] = b.Active
	}

	if !blockMap["session"] {
		t.Error("session should be active when ActiveBackends.Session is non-nil")
	}
	if blockMap["cache"] {
		t.Error("cache should be inactive when ActiveBackends.Cache is nil")
	}
}
