//ff:func feature=gen-react type=test control=sequence
//ff:what writeLibArtifacts — utils.ts 상시 방출 + sitemap 유무에 따른 breadcrumbs.ts 방출/미방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestWriteLibArtifacts(t *testing.T) {
	t.Run("without a sitemap only utils.ts appears", func(t *testing.T) {
		srcDir := t.TempDir()
		if err := writeLibArtifacts(srcDir, &yongol.Fullstack{}, nil); err != nil {
			t.Fatalf("writeLibArtifacts: %v", err)
		}
		if _, err := os.Stat(filepath.Join(srcDir, "lib", "utils.ts")); err != nil {
			t.Errorf("utils.ts missing: %v", err)
		}
		if _, err := os.Stat(filepath.Join(srcDir, "lib", "breadcrumbs.ts")); !os.IsNotExist(err) {
			t.Error("breadcrumbs.ts must not exist without a sitemap")
		}
	})

	t.Run("with a sitemap the breadcrumb pair appears too", func(t *testing.T) {
		srcDir := t.TempDir()
		fs := &yongol.Fullstack{Sitemap: &stml.SitemapSpec{}}
		if err := writeLibArtifacts(srcDir, fs, nil); err != nil {
			t.Fatalf("writeLibArtifacts: %v", err)
		}
		if _, err := os.Stat(filepath.Join(srcDir, "lib", "breadcrumbs.ts")); err != nil {
			t.Errorf("breadcrumbs.ts missing: %v", err)
		}
		if _, err := os.Stat(filepath.Join(srcDir, "components", "ui", "Breadcrumb.tsx")); err != nil {
			t.Errorf("Breadcrumb.tsx missing: %v", err)
		}
	})
}
