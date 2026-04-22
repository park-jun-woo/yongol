//ff:func feature=chain type=util control=iteration dimension=2
//ff:what callRef에 매칭되는 FuncSpec을 찾아 Link를 반환한다
package chain

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// findFuncSpecLink searches specs for a matching func spec and returns a Link.
func findFuncSpecLink(callRef, pkg, funcName string, specs []funcspec.FuncSpec, specsDir string) (Link, bool) {
	for _, spec := range specs {
		if spec.Package != pkg || !strings.EqualFold(spec.Name, funcName) {
			continue
		}
		relPath := resolveFuncSpecPath(spec, funcName, specsDir)
		line := grepLine(filepath.Join(specsDir, relPath), funcName)
		return Link{
			Kind:    "FuncSpec",
			File:    relPath,
			Line:    line,
			Summary: "@func " + callRef,
		}, true
	}
	return Link{}, false
}
