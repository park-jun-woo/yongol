//ff:func feature=gen-gogin type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 gogin/ssac.Generate 통합 커버리지
package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGoginSSaCGenerate_Integ(t *testing.T) {
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

	serviceDir := filepath.Join(out, "backend", "internal", "service")
	if _, err := os.Stat(filepath.Join(serviceDir, "server.go")); err != nil {
		t.Errorf("expected server.go: %v", err)
	}
	var goFiles int
	_ = filepath.Walk(serviceDir, func(p string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() && filepath.Ext(p) == ".go" {
			goFiles++
		}
		return nil
	})
	if goFiles < 2 {
		t.Errorf("expected multiple generated .go service files, got %d", goFiles)
	}
}
