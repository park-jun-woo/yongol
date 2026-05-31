//ff:func feature=gen-react type=test control=sequence
//ff:what TestGenerateFrontendSetup_ZeroCov — generateFrontendSetup 을 빈 Fullstack 으로 직접 호출

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateFrontendSetup_ZeroCov(t *testing.T) {
	out := t.TempDir()
	// Empty Fullstack drives the nil-guard branches (no Manifest/DesignSpec).
	if err := generateFrontendSetup(&yongol.Fullstack{}, out); err != nil {
		t.Fatalf("generateFrontendSetup: %v", err)
	}
	for _, name := range []string{
		"package.json",
		"vite.config.ts",
		filepath.Join("src", "main.tsx"),
	} {
		if _, err := os.Stat(filepath.Join(out, "frontend", name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}
