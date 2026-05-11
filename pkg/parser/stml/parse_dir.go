//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 디렉토리 내 모든 .html 파일을 파싱하여 PageSpec 목록 반환
package stml

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir parses all .html files in the given directory and returns a PageSpec for each.
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
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".html") {
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
