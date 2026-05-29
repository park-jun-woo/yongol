//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildInfraPlanSession -- session postgres 활성 → SessionConfig 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildInfraPlanSession(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Module: "github.com/test/sess"},
		},
	}
	ps := &prepared.State{
		ActiveBackends: prepared.ActiveBackends{
			Session: &prepared.Session{Backend: "postgres"},
		},
	}

	plan := BuildInfraPlan(fs, ps)

	if plan.Session == nil {
		t.Fatal("Session should be non-nil")
	}
	if plan.Session.Backend != "postgres" {
		t.Errorf("Session.Backend = %q, want %q", plan.Session.Backend, "postgres")
	}
}
