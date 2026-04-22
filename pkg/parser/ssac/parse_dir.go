//ff:func feature=ssac-parse type=parser control=sequence
//ff:what ParseDir — recursively walks a directory for all .ssac files and returns []ServiceFunc
package ssac

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir recursively walks a directory for all .ssac files and returns the extracted []ServiceFunc.
func ParseDir(dir string) ([]ServiceFunc, []diagnostic.Diagnostic) {
	var funcs []ServiceFunc
	var diags []diagnostic.Diagnostic
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    0,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "directory walk error: " + err.Error(),
			})
			return nil
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".ssac") {
			return nil
		}
		parsed, d2 := parseDirEntry(dir, path, d.Name())
		diags = append(diags, d2...)
		funcs = append(funcs, parsed...)
		return nil
	})
	return funcs, diags
}
