//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWritePrismaModule_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writePrismaModule(dir); err != nil {
		t.Fatalf("writePrismaModule error: %v", err)
	}
	for _, name := range []string{"prisma/prisma.service.ts", "prisma/prisma.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}
