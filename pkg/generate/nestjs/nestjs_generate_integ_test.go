//ff:func feature=gen-nestjs type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 nestjs.Generate 통합 커버리지
package nestjs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNestjsGenerate_Integ(t *testing.T) {
	root := "/home/parkjunwoo/.clari/repos/fullend/dummys/zenflow/try-03/specs"
	if _, err := os.Stat(root); err != nil {
		t.Skipf("dummy specs not present: %v", err)
	}
	fs := loadDummyFS_Integ(t, root)
	if len(fs.ServiceFuncs) == 0 {
		t.Fatal("expected ServiceFuncs from dummy specs")
	}

	out := t.TempDir()
	if err := Generate(fs, out); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	// Feature module + service artifacts should exist somewhere under the tree.
	var tsFiles int
	_ = filepath.Walk(out, func(p string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() && filepath.Ext(p) == ".ts" {
			tsFiles++
		}
		return nil
	})
	if tsFiles == 0 {
		t.Error("expected generated .ts files under output dir")
	}
}
