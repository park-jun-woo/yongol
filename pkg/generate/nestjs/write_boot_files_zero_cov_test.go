//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteBootFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	plan := &ir.BootPlan{ProjectID: "myapp"}
	if err := writeBootFiles(dir, plan, []string{"users"}, []string{"queue"}); err != nil {
		t.Fatalf("writeBootFiles error: %v", err)
	}
	for _, name := range []string{"main.ts", "app.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}
