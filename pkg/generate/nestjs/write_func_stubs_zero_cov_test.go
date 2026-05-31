//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFuncStubs_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	pkgs := []externalPackage{{Name: "billing", Methods: []string{"Charge"}}}
	if err := writeFuncStubs(dir, pkgs); err != nil {
		t.Fatalf("writeFuncStubs error: %v", err)
	}
	for _, name := range []string{"billing/billing.service.ts", "billing/billing.module.ts"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}
