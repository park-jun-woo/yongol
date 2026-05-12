//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what ParseLayoutDir — 레이아웃 디렉토리의 .html 파일들을 순회하며 LayoutSpec 목록 반환

package stml

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseLayoutDir parses all .html files in the given layouts directory
// and returns a LayoutSpec for each.
func ParseLayoutDir(dir string) ([]LayoutSpec, []diagnostic.Diagnostic) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("read layout dir %s: %s", dir, err),
		}}
	}

	var layouts []LayoutSpec
	var allDiags []diagnostic.Diagnostic
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".html") {
			continue
		}
		layout, diags := ParseLayoutFile(filepath.Join(dir, e.Name()))
		if len(diags) > 0 {
			allDiags = append(allDiags, diags...)
			continue
		}
		layouts = append(layouts, layout)
	}
	if len(allDiags) > 0 {
		return nil, allDiags
	}
	return layouts, nil
}
