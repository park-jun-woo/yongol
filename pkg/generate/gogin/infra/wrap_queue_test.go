//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitQueueWrapper — queue adapter emit + 누락 포트/write 에러 검증

package infra

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func TestEmitQueueWrapper(t *testing.T) {
	iface := &ssacmeta.PackageInterface{Package: "queue"}
	ports := []ssacmeta.Port{{Name: "QueuePublish"}}

	t.Run("MissingPublishError", func(t *testing.T) {
		err := emitQueueWrapper(iface, []ssacmeta.Port{{Name: "Other"}}, "example.com/app", t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "missing QueuePublish") {
			t.Errorf("expected missing QueuePublish error, got: %v", err)
		}
	})

	t.Run("EmitsPostgresGo", func(t *testing.T) {
		arts := t.TempDir()
		if err := emitQueueWrapper(iface, ports, "example.com/app", arts); err != nil {
			t.Fatalf("emitQueueWrapper error: %v", err)
		}
		path := filepath.Join(arts, "backend", "internal", "infra", "queue", "postgres.go")
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected postgres.go: %v", err)
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		if err := os.WriteFile(filepath.Join(arts, "backend"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := emitQueueWrapper(iface, ports, "example.com/app", arts); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}
