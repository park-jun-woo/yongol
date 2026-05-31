//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateGeneratesClaimAndBearerAuth — user_claim.go + bearerauth.go 산출 검증
package auth

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateGeneratesClaimAndBearerAuth(t *testing.T) {
	dir := t.TempDir()
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Module: "github.com/test/app"},
		},
	}
	p := prepared.State{Auth: prepared.Auth{
		Present: true,
		Mode:    "bearer",
		Raw: &manifest.Auth{
			Claims: map[string]manifest.ClaimDef{
				"ID":    {Key: "user_id", GoType: "int64", Typed: true},
				"Email": {Key: "email", GoType: "string"},
			},
		},
	}}
	if err := Generate(fs, p, dir); err != nil {
		t.Fatalf("Generate error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "backend", "internal", "model", "user_claim.go")); err != nil {
		t.Errorf("expected user_claim.go emitted: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "backend", "internal", "middleware", "bearerauth.go")); err != nil {
		t.Errorf("expected bearerauth.go emitted: %v", err)
	}
}
