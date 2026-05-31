//ff:func feature=generate type=test control=sequence
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMakeFrontendCopyWalker_ZeroCov(t *testing.T) {
	srcRoot := t.TempDir()
	dstRoot := t.TempDir()
	// create: a copied file, a managed file (skipped), a non-copied ext (skipped),
	// node_modules dir (skipped).
	mustWrite := func(rel, content string) string {
		p := filepath.Join(srcRoot, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}
	plain := mustWrite("src/pages/Home.tsx", "x")
	managed := mustWrite("src/api.ts", "y")
	other := mustWrite("src/readme.md", "z")
	nm := mustWrite("node_modules/pkg/index.js", "n")

	walker := makeFrontendCopyWalker(srcRoot, dstRoot)

	// walkErr propagation.
	if err := walker("p", nil, errTestWalk); err != errTestWalk {
		t.Errorf("walkErr not propagated: %v", err)
	}

	// node_modules dir → SkipDir.
	nmDir := filepath.Join(srcRoot, "node_modules")
	fi, _ := os.Stat(nmDir)
	if err := walker(nmDir, fi, nil); err != filepath.SkipDir {
		t.Errorf("node_modules should SkipDir, got %v", err)
	}
	// non-node_modules dir → nil.
	pagesFi, _ := os.Stat(filepath.Join(srcRoot, "src", "pages"))
	if err := walker(filepath.Join(srcRoot, "src", "pages"), pagesFi, nil); err != nil {
		t.Errorf("dir walk = %v", err)
	}

	walkFile := func(p string) error {
		fi, err := os.Stat(p)
		if err != nil {
			t.Fatal(err)
		}
		return walker(p, fi, nil)
	}
	if err := walkFile(plain); err != nil {
		t.Errorf("plain copy err: %v", err)
	}
	if err := walkFile(managed); err != nil {
		t.Errorf("managed err: %v", err)
	}
	if err := walkFile(other); err != nil {
		t.Errorf("other err: %v", err)
	}
	_ = nm

	// plain should be copied; managed and other should not.
	if _, err := os.Stat(filepath.Join(dstRoot, "src/pages/Home.tsx")); err != nil {
		t.Errorf("plain not copied: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dstRoot, "src/api.ts")); !os.IsNotExist(err) {
		t.Error("managed file should not be copied")
	}
	if _, err := os.Stat(filepath.Join(dstRoot, "src/readme.md")); !os.IsNotExist(err) {
		t.Error("non-copied ext should not be copied")
	}
}
