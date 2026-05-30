//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitSessionWrapper — session adapter emit + 누락 포트/write 에러 검증

package infra

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func sessionPorts() []ssacmeta.Port {
	return []ssacmeta.Port{
		{Name: "SessionSet"},
		{Name: "SessionGet"},
		{Name: "SessionDelete"},
	}
}

func TestEmitSessionWrapper(t *testing.T) {
	iface := &ssacmeta.PackageInterface{Package: "session"}

	t.Run("MissingPortsError", func(t *testing.T) {
		err := emitSessionWrapper(iface, []ssacmeta.Port{{Name: "SessionSet"}}, "example.com/app", t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "missing one of") {
			t.Errorf("expected missing-port error, got: %v", err)
		}
	})

	t.Run("EmitsPostgresGo", func(t *testing.T) {
		arts := t.TempDir()
		if err := emitSessionWrapper(iface, sessionPorts(), "example.com/app", arts); err != nil {
			t.Fatalf("emitSessionWrapper error: %v", err)
		}
		path := filepath.Join(arts, "backend", "internal", "infra", "session", "postgres.go")
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected postgres.go: %v", err)
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		if err := os.WriteFile(filepath.Join(arts, "backend"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := emitSessionWrapper(iface, sessionPorts(), "example.com/app", arts); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}
