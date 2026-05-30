//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitAll — SsacInterfaces 순회 emit: nil/empty/loop(skip) 경로 검증

package infra

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestEmitAll(t *testing.T) {
	t.Run("NilFS", func(t *testing.T) {
		if err := EmitAll(nil, t.TempDir(), "example.com/app"); err != nil {
			t.Errorf("expected nil for nil fs, got: %v", err)
		}
	})

	t.Run("NoInterfaces", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		if err := EmitAll(fs, t.TempDir(), "example.com/app"); err != nil {
			t.Errorf("expected nil for no interfaces, got: %v", err)
		}
	})

	t.Run("EmitErrorWrapped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SsacInterfaces: map[string]*ssacmeta.PackageInterface{
				"cache": {Package: "cache", Ports: []ssacmeta.Port{{Name: "Bogus", When: "always"}}},
			},
		}
		err := EmitAll(fs, t.TempDir(), "example.com/app")
		if err == nil || !strings.Contains(err.Error(), "infra: emit cache") {
			t.Errorf("expected wrapped emit error, got: %v", err)
		}
	})

	t.Run("IteratesAndSkipsUnregistered", func(t *testing.T) {
		arts := t.TempDir()
		// Two packages with active ports but no registered emitter -> each
		// emitOnePackage returns nil; EmitAll walks both in sorted order.
		fs := &yongol.Fullstack{
			SsacInterfaces: map[string]*ssacmeta.PackageInterface{
				"zeta":  {Package: "zeta", Ports: []ssacmeta.Port{{Name: "p", When: "always"}}},
				"alpha": {Package: "alpha", Ports: []ssacmeta.Port{{Name: "p", When: "always"}}},
			},
		}
		if err := EmitAll(fs, arts, "example.com/app"); err != nil {
			t.Fatalf("EmitAll error: %v", err)
		}
		// Unregistered emitters produce no files.
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "infra")); !os.IsNotExist(err) {
			t.Errorf("expected no infra output for unregistered packages, stat err: %v", err)
		}
	})
}
