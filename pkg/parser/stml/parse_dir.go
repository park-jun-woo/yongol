//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 디렉토리 내 모든 .html 파일을 파싱하여 PageSpec 목록 반환 (sitemap.html 은 페이지가 아니므로 제외)
package stml

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir parses all .html files in the given directory and returns a
// PageSpec for each. sitemap.html is the site-structure declaration, not a
// page — it is parsed separately by ParseSitemap (otherwise it would be
// mis-parsed as a page mounted at /sitemap).
func ParseDir(dir string) ([]PageSpec, []diagnostic.Diagnostic) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("read dir %s: %s", dir, err),
		}}
	}

	var pages []PageSpec
	var allDiags []diagnostic.Diagnostic
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".html") || e.Name() == "sitemap.html" {
			continue
		}
		page, diags := ParseFile(filepath.Join(dir, e.Name()))
		if len(diags) > 0 {
			allDiags = append(allDiags, diags...)
			continue
		}
		pages = append(pages, page)
	}
	if len(allDiags) > 0 {
		return nil, allDiags
	}
	return pages, nil
}
