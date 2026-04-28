//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestLoadPackageInterfacesWalksPkgDir — pkg/*/interface.yaml 를 모두 수집

package ssacmeta

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPackageInterfacesWalksPkgDir(t *testing.T) {
	root := t.TempDir()
	pkgDir := filepath.Join(root, "pkg")
	for _, name := range []string{"cache", "session", "notAnSsacPkg"} {
		if err := os.MkdirAll(filepath.Join(pkgDir, name), 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
	}
	writePkgInterfaceFixture(t, pkgDir, "cache", "cache")
	writePkgInterfaceFixture(t, pkgDir, "session", "session")
	// notAnSsacPkg intentionally has no interface.yaml — should be skipped.

	ifaces, err := LoadPackageInterfaces(root)
	if err != nil {
		t.Fatalf("LoadPackageInterfaces err: %v", err)
	}
	if _, ok := ifaces["cache"]; !ok {
		t.Errorf("cache not loaded")
	}
	if _, ok := ifaces["session"]; !ok {
		t.Errorf("session not loaded")
	}
	if _, ok := ifaces["notAnSsacPkg"]; ok {
		t.Errorf("non-ssac pkg must be skipped")
	}
}
