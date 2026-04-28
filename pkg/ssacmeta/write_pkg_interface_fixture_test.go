//ff:func feature=ssacmeta type=test-helper control=sequence
//ff:what writePkgInterfaceFixture — 테스트용 interface.yaml 한 건 기록

package ssacmeta

import (
	"os"
	"path/filepath"
	"testing"
)

// writePkgInterfaceFixture writes a minimal interface.yaml under
// pkgDir/<name>/interface.yaml declaring `package: <pkg>` and zero ports.
// Used by LoadPackageInterfaces walk tests to populate fixture trees.
func writePkgInterfaceFixture(t *testing.T, pkgDir, name, pkg string) {
	t.Helper()
	src := "version: 1\npackage: " + pkg + "\nports: []\n"
	if err := os.WriteFile(filepath.Join(pkgDir, name, "interface.yaml"), []byte(src), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
}
