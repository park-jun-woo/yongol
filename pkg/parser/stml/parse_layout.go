//ff:func feature=stml-parse type=parser control=sequence
//ff:what ParseLayoutFile — 단일 레이아웃 HTML 파일을 읽어 LayoutSpec 반환

package stml

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseLayoutFile parses a single layout HTML file and returns a LayoutSpec.
func ParseLayoutFile(path string) (LayoutSpec, []diagnostic.Diagnostic) {
	f, err := os.Open(path)
	if err != nil {
		return LayoutSpec{}, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: err.Error(),
		}}
	}
	defer f.Close()

	return ParseLayoutReader(filepath.Base(path), path, f)
}
