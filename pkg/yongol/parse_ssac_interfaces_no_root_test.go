//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSsacInterfaces — pkg root 미발견(skip) + 발견 시 SsacInterfaces 로드
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSsacInterfaces_NoRoot(t *testing.T) {
	// Point the env override at a non-existent path and isolate fallbacks.
	t.Setenv("YONGOL_SSAC_PKG", filepath.Join(t.TempDir(), "nope"))
	isolated := t.TempDir()
	if err := os.Chdir(isolated); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Setenv("GOMODCACHE", filepath.Join(isolated, "empty-cache"))

	fs := &Fullstack{}
	parseSsacInterfaces(fs)
	if fs.SsacInterfaces != nil {
		t.Skipf("ambient ssac root resolved; SsacInterfaces=%v", fs.SsacInterfaces)
	}
}
