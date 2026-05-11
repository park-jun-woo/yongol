//ff:func feature=stml-parse type=parser control=sequence
//ff:what 단일 HTML 파일을 파싱하여 PageSpec 반환
package stml

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseFile parses a single HTML file and returns a PageSpec.
func ParseFile(path string) (PageSpec, []diagnostic.Diagnostic) {
	f, err := os.Open(path)
	if err != nil {
		return PageSpec{}, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: err.Error(),
		}}
	}
	defer f.Close()

	return ParseReader(filepath.Base(path), f)
}
