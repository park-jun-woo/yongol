//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what collectOverrides — STML 파일에서 <!-- @override class="..." --> 주석의 class 값을 수집
package stml_design

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectOverrides scans raw STML HTML files for <!-- @override class="..." --> comments
// and returns the set of class attribute values extracted from those comments.
func collectOverrides(fs *yongol.Fullstack) overrideSet {
	result := make(overrideSet)
	frontendDir := filepath.Join(fs.SpecsDir, "frontend")

	for _, page := range fs.STMLPages {
		path := filepath.Join(frontendDir, page.FileName)
		classes := parseOverridesFromFile(path)
		if len(classes) > 0 {
			result[page.FileName] = classes
		}
	}
	return result
}
