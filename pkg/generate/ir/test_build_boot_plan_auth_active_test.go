//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestBuildBootPlanAuthActive -- auth 선언 시 jwt-secret + auth-init 블록 활성화 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildBootPlanAuthActive(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Metadata: manifest.Metadata{Name: "authproj"},
			Backend: manifest.Backend{
				Module: "github.com/test/authproj",
				Auth: &manifest.Auth{
					Claims: map[string]manifest.ClaimDef{
						"UserID": {Key: "user_id"},
					},
				},
			},
		},
	}
	ps := &prepared.State{
		Auth: prepared.Auth{
			Present: true,
			Mode:    "bearer",
			Raw:     fs.Manifest.Backend.Auth,
		},
	}

	plan := BuildBootPlan(fs, ps, "go-gin")

	blockMap := map[string]bool{}
	for _, b := range plan.ActiveBlocks {
		blockMap[b.Name] = b.Active
	}

	if !blockMap["jwt-secret"] {
		t.Error("jwt-secret should be active when auth has claims")
	}
	if !blockMap["auth-init"] {
		t.Error("auth-init should be active when auth is present")
	}
}
