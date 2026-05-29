//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildInfraPlanAllActive -- 전체 인프라 활성 시 모든 필드 non-nil 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildInfraPlanAllActive(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module: "github.com/test/all",
				Auth: &manifest.Auth{
					Claims: map[string]manifest.ClaimDef{"UserID": {Key: "uid"}},
				},
			},
		},
	}
	ps := &prepared.State{
		ActiveBackends: prepared.ActiveBackends{
			Session: &prepared.Session{Backend: "postgres"},
			Cache:   &prepared.Cache{Backend: "postgres"},
			Queue:   &prepared.Queue{Backend: "postgres"},
		},
		Auth: prepared.Auth{
			Present: true,
			Mode:    "bearer",
			Raw:     fs.Manifest.Backend.Auth,
		},
	}

	plan := BuildInfraPlan(fs, ps)

	if plan.Session == nil {
		t.Error("Session should be non-nil")
	}
	if plan.Cache == nil {
		t.Error("Cache should be non-nil")
	}
	if plan.Queue == nil {
		t.Error("Queue should be non-nil")
	}
	if plan.Auth == nil {
		t.Error("Auth should be non-nil")
	}
}
