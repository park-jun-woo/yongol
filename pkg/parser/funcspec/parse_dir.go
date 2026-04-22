//ff:func feature=funcspec type=parser control=sequence
//ff:what 디렉토리 내 모든 .go 파일을 재귀적으로 파싱하여 FuncSpec 목록을 반환한다
package funcspec

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir parses all .go files under dir (recursively by package subdirectory).
// Returns a flat list of FuncSpecs and accumulated diagnostics.
func ParseDir(dir string) ([]FuncSpec, []diagnostic.Diagnostic) {
	var specs []FuncSpec
	var specDirs []string
	var diags []diagnostic.Diagnostic

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    0,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "walk error: " + err.Error(),
			})
			return nil
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			return nil
		}
		fs, dd := ParseFile(path)
		if len(dd) > 0 {
			diags = append(diags, dd...)
			return nil
		}
		if fs != nil {
			// Derive package from parent dir name.
			rel, _ := filepath.Rel(dir, path)
			parts := strings.Split(filepath.Dir(rel), string(filepath.Separator))
			if parts[0] != "." {
				fs.Package = parts[0]
			}
			specs = append(specs, *fs)
			specDirs = append(specDirs, filepath.Dir(path))
		}
		return nil
	})

	diags = append(diags, fillMissingFields(specs, specDirs)...)
	return specs, diags
}
