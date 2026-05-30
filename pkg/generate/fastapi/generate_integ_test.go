//ff:func feature=gen-fastapi type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 fastapi.Generate 통합 커버리지

package fastapi

import (
	"os"
	"path/filepath"
	"testing"

	pssac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func loadDummyFS_Integ(t *testing.T, root string) *yongol.Fullstack {
	t.Helper()
	det, err := yongol.DetectSSOTs(root)
	if err != nil {
		t.Fatalf("DetectSSOTs(%s): %v", root, err)
	}
	fs := yongol.ParseAll(root, det)
	if len(fs.ServiceFuncs) == 0 {
		funcs, _ := pssac.ParseDir(filepath.Join(root, "service"))
		fs.ServiceFuncs = funcs
	}
	return fs
}

// TestFastapiGenerate_Integ runs the full fastapi.Generate pipeline against the
// zenflow dummy specs, materializing the backend/app tree (scaffold, models,
// schemas, feature modules, boot files).
func TestFastapiGenerate_Integ(t *testing.T) {
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
