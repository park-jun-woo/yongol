//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSsacInterfaces — pkg root 미발견(skip) + 발견 시 SsacInterfaces 로드
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSsacInterfaces_Loaded(t *testing.T) {
	base := t.TempDir()
	pkgDir := filepath.Join(base, "pkg")
	cacheDir := filepath.Join(pkgDir, "cache")
	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cacheDir, "interface.yaml"),
		[]byte("version: 1\npackage: cache\nports: []\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// findYongolPkgRoot returns YONGOL_SSAC_PKG when it is an existing dir.
	t.Setenv("YONGOL_SSAC_PKG", pkgDir)

	fs := &Fullstack{}
	parseSsacInterfaces(fs)
	if fs.SsacInterfaces == nil {
		t.Fatalf("expected SsacInterfaces to be loaded")
	}
	if _, ok := fs.SsacInterfaces["cache"]; !ok {
		t.Fatalf("expected 'cache' package interface, got %+v", fs.SsacInterfaces)
	}
}
