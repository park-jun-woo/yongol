//ff:func feature=chain type=util control=iteration dimension=1
//ff:what resolveFuncSpecPath finds the actual file path for a func spec
package chain

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// resolveFuncSpecPath finds the actual file path for a func spec.
func resolveFuncSpecPath(spec funcspec.FuncSpec, funcName, specsDir string) string {
	relPath := "func/" + spec.Package + "/" + toSnakeCase(spec.Name) + ".go"
	if _, err := os.Stat(filepath.Join(specsDir, relPath)); err == nil {
		return relPath
	}
	matches, _ := filepath.Glob(filepath.Join(specsDir, "func", spec.Package, "*.go"))
	for _, m := range matches {
		if grepLine(m, "@func") > 0 && grepLine(m, funcName) > 0 {
			rel, _ := filepath.Rel(specsDir, m)
			return rel
		}
	}
	return relPath
}
