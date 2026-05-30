//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteBootFiles — main.py + config.py + database.py 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteBootFiles(t *testing.T) {
	t.Run("WritesAllBootFiles", func(t *testing.T) {
		appDir := t.TempDir()
		plan := &ir.BootPlan{ProjectID: "myproject"}
		if err := writeBootFiles(appDir, plan, []string{"users"}); err != nil {
			t.Fatalf("writeBootFiles error: %v", err)
		}
		for _, name := range []string{"main.py", "config.py", "database.py"} {
			info, err := os.Stat(filepath.Join(appDir, name))
			if err != nil {
				t.Errorf("expected %s: %v", name, err)
				continue
			}
			if info.Size() == 0 {
				t.Errorf("%s is empty", name)
			}
		}
	})

	t.Run("RenderMainFailsOnNilPlan", func(t *testing.T) {
		appDir := t.TempDir()
		if err := writeBootFiles(appDir, nil, nil); err == nil {
			t.Errorf("expected RenderMain error for nil plan, got nil")
		}
	})

	t.Run("WriteMainFailsOnBadDir", func(t *testing.T) {
		base := t.TempDir()
		filePath := filepath.Join(base, "file")
		if err := os.WriteFile(filePath, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// appDir is a regular file -> WriteFile(main.py) fails.
		plan := &ir.BootPlan{ProjectID: "p"}
		if err := writeBootFiles(filePath, plan, nil); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}
