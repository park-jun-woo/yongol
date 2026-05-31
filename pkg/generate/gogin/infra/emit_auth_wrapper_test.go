//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitAuthWrapper — auth RefreshStore adapter 6파일 emit + 누락/mkdir/write 에러 검증
package infra

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func TestEmitAuthWrapper(t *testing.T) {
	iface := &ssacmeta.PackageInterface{Package: "auth"}

	t.Run("MissingPortsError", func(t *testing.T) {
		err := emitAuthWrapper(iface, []ssacmeta.Port{{Name: "RefreshTokenInsert"}}, "example.com/app", t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "missing one of") {
			t.Errorf("expected missing-port error, got: %v", err)
		}
	})

	t.Run("EmitsSixFiles", func(t *testing.T) {
		arts := t.TempDir()
		if err := emitAuthWrapper(iface, authPorts(), "example.com/app", arts); err != nil {
			t.Fatalf("emitAuthWrapper error: %v", err)
		}
		dir := filepath.Join(arts, "backend", "internal", "infra", "auth")
		for _, name := range []string{
			"postgres.go", "postgres_new.go", "postgres_create.go",
			"postgres_consume.go", "postgres_revoke.go", "postgres_revoke_all.go",
		} {
			if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
				t.Errorf("expected %s: %v", name, err)
			}
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		arts := t.TempDir()
		// infra parent collides with a file.
		infraParent := filepath.Join(arts, "backend", "internal", "infra")
		if err := os.MkdirAll(filepath.Dir(infraParent), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(infraParent, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := emitAuthWrapper(iface, authPorts(), "example.com/app", arts); err == nil {
			t.Errorf("expected mkdir error, got nil")
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		dir := filepath.Join(arts, "backend", "internal", "infra", "auth")
		// Pre-create every target as a directory so WriteFile fails.
		for _, name := range []string{
			"postgres.go", "postgres_new.go", "postgres_create.go",
			"postgres_consume.go", "postgres_revoke.go", "postgres_revoke_all.go",
		} {
			if err := os.MkdirAll(filepath.Join(dir, name), 0o755); err != nil {
				t.Fatalf("setup: %v", err)
			}
		}
		err := emitAuthWrapper(iface, authPorts(), "example.com/app", arts)
		if err == nil || !strings.Contains(err.Error(), "write ") {
			t.Errorf("expected write error, got: %v", err)
		}
	})
}
