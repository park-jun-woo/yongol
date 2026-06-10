//ff:func feature=gen-fastapi type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 fastapi.Generate 통합 커버리지
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFastapiGenerate_Integ(t *testing.T) {
	root := "/home/parkjunwoo/.clari/repos/fullend/examples/zenflow/try-03/specs"
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

	appDir := filepath.Join(out, "backend", "app")
	if _, err := os.Stat(filepath.Join(appDir, "main.py")); err != nil {
		t.Errorf("expected main.py: %v", err)
	}
	var pyFiles int
	_ = filepath.Walk(appDir, func(p string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() && filepath.Ext(p) == ".py" {
			pyFiles++
		}
		return nil
	})
	if pyFiles < 3 {
		t.Errorf("expected several generated .py files, got %d", pyFiles)
	}
}
