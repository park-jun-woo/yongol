//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSsacInterfaces — pkg root 미발견(skip) + 발견 시 SsacInterfaces 로드
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSsacInterfaces_LoadError(t *testing.T) {
	base := t.TempDir()
	pkgDir := filepath.Join(base, "pkg")
	badDir := filepath.Join(pkgDir, "broken")
	if err := os.MkdirAll(badDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// Malformed YAML → LoadPackageInterface returns an error → warn + return.
	if err := os.WriteFile(filepath.Join(badDir, "interface.yaml"),
		[]byte("version: 1\npackage: [this is: not valid\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("YONGOL_SSAC_PKG", pkgDir)

	fs := &Fullstack{}
	parseSsacInterfaces(fs)
	if fs.SsacInterfaces != nil {
		t.Fatalf("expected nil SsacInterfaces on load error, got %+v", fs.SsacInterfaces)
	}
}
