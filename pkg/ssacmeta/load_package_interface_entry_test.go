//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestloadPackageInterfaceEntry — loadPackageInterfaceEntry() dir/파일/키 폴백 분기

package ssacmeta

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPackageInterfaceEntry(t *testing.T) {
	pkgRoot := t.TempDir()

	// 1) dir with interface.yaml carrying an explicit package: key.
	mustMkdir(t, filepath.Join(pkgRoot, "cachedir"))
	mustWrite(t, filepath.Join(pkgRoot, "cachedir", "interface.yaml"),
		"version: 1\npackage: cache\nports: []\n")

	// 2) dir with interface.yaml but no package: key -> falls back to dir name.
	mustMkdir(t, filepath.Join(pkgRoot, "session"))
	mustWrite(t, filepath.Join(pkgRoot, "session", "interface.yaml"),
		"version: 1\nports: []\n")

	// 3) dir without any interface.yaml -> skipped.
	mustMkdir(t, filepath.Join(pkgRoot, "empty"))

	// 4) a plain file (non-dir) -> skipped.
	mustWrite(t, filepath.Join(pkgRoot, "README.md"), "not a package\n")

	entries, err := os.ReadDir(pkgRoot)
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}

	out := map[string]*PackageInterface{}
	for _, e := range entries {
		if err := loadPackageInterfaceEntry(out, pkgRoot, e); err != nil {
			t.Fatalf("loadPackageInterfaceEntry(%s): %v", e.Name(), err)
		}
	}

	if _, ok := out["cache"]; !ok {
		t.Errorf("expected key from package: field (cache)")
	}
	if _, ok := out["session"]; !ok {
		t.Errorf("expected fallback key from dir name (session)")
	}
	if _, ok := out["empty"]; ok {
		t.Errorf("dir without interface.yaml must be skipped")
	}
	if _, ok := out["README.md"]; ok {
		t.Errorf("non-dir entry must be skipped")
	}
	if len(out) != 2 {
		t.Errorf("out size = %d, want 2 (%v)", len(out), keysOf(out))
	}
}

func mustMkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func keysOf(m map[string]*PackageInterface) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
