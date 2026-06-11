//ff:func feature=stml-parse type=parser control=sequence
//ff:what ParseSitemap — frontend/sitemap.html 파일을 열어 SitemapSpec 반환

package stml

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseSitemap parses the fixed-name frontend/sitemap.html file and returns
// a SitemapSpec (plans/stml/sitemap Phase001).
func ParseSitemap(path string) (SitemapSpec, []diagnostic.Diagnostic) {
	f, err := os.Open(path)
	if err != nil {
		return SitemapSpec{}, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: err.Error(),
		}}
	}
	defer f.Close()

	return ParseSitemapReader(filepath.Base(path), f)
}
