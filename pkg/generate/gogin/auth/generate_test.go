//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — auth UserClaim+BearerAuth 생성 skip/success 경로 검증

package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate(t *testing.T) {
	t.Run("SkipsWhenNoAuth", func(t *testing.T) {
		dir := t.TempDir()
		fs := &yongol.Fullstack{}
		p := prepared.State{Auth: prepared.Auth{Present: false}}
		if err := Generate(fs, p, dir); err != nil {
			t.Fatalf("Generate error: %v", err)
		}
		// Nothing should be emitted.
		if _, err := os.Stat(filepath.Join(dir, "backend")); !os.IsNotExist(err) {
			t.Errorf("expected no output for absent auth, stat err: %v", err)
		}
	})

	t.Run("SkipsWhenEmptyClaims", func(t *testing.T) {
		dir := t.TempDir()
		fs := &yongol.Fullstack{}
		p := prepared.State{Auth: prepared.Auth{
			Present: true,
			Raw:     &manifest.Auth{Claims: map[string]manifest.ClaimDef{}},
		}}
		if err := Generate(fs, p, dir); err != nil {
			t.Fatalf("Generate error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "backend")); !os.IsNotExist(err) {
			t.Errorf("expected no output for empty claims, stat err: %v", err)
		}
	})

	claimsState := func() prepared.State {
		return prepared.State{Auth: prepared.Auth{
			Present: true,
			Mode:    "bearer",
			Raw: &manifest.Auth{
				Claims: map[string]manifest.ClaimDef{
					"ID": {Key: "user_id", GoType: "int64", Typed: true},
				},
			},
		}}
	}
	fsWithModule := func() *yongol.Fullstack {
		return &yongol.Fullstack{
			Manifest: &manifest.ProjectConfig{
				Backend: manifest.Backend{Module: "github.com/test/app"},
			},
		}
	}

	t.Run("UserClaimError", func(t *testing.T) {
		dir := t.TempDir()
		// Make backend/internal a regular file so model dir MkdirAll fails.
		internal := filepath.Join(dir, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := Generate(fsWithModule(), claimsState(), dir)
		if err == nil || !strings.Contains(err.Error(), "user_claim") {
			t.Errorf("expected user_claim error, got: %v", err)
		}
	})

	t.Run("BearerAuthError", func(t *testing.T) {
		dir := t.TempDir()
		// model dir creatable, but middleware path collides with a file so
		// generateBearerAuth's MkdirAll fails after user_claim succeeds.
		mwParent := filepath.Join(dir, "backend", "internal")
		if err := os.MkdirAll(mwParent, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(filepath.Join(mwParent, "middleware"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := Generate(fsWithModule(), claimsState(), dir)
		if err == nil || !strings.Contains(err.Error(), "bearer_auth") {
			t.Errorf("expected bearer_auth error, got: %v", err)
		}
	})

	t.Run("GeneratesClaimAndBearerAuth", func(t *testing.T) {
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
	})
}
