//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what collectPackageTypes — collects struct types and fields from all .go files in a directory; reports Go parse failures as Diagnostics
package funcspec

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// collectPackageTypes parses all .go files in dir (non-recursive)
// and returns a map of struct name to fields.
//
// Diagnostic contract:
//   - If dir does not exist, returns SILENT-OK (absent companion type directory is normal).
//   - Any other ReadDir error returns one Diagnostic + empty result.
//   - A parser.ParseFile failure on an individual file is recorded as one Diagnostic
//     while the remaining files continue to be collected (partial success).
func collectPackageTypes(dir string) (map[string][]Field, []diagnostic.Diagnostic) {
	result := make(map[string][]Field)
	var diags []diagnostic.Diagnostic

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return result, nil
		}
		return result, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "cannot read funcspec type dir: " + err.Error(),
		}}
	}

	fset := token.NewFileSet()
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    extractGoParseErrorLine(err),
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "Go parse failed: " + err.Error(),
			})
			continue
		}
		collectStructsFromFile(f, result)
	}
	return result, diags
}
