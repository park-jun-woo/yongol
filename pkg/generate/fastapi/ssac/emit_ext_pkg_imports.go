//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what emitExtPkgImports — 외부 패키지 함수 import 출력

package ssac

import (
	"fmt"
	"sort"
	"strings"
)

// emitExtPkgImports writes external package function imports.
func emitExtPkgImports(b *strings.Builder, extPkgs map[string]map[string]bool) {
	pkgNames := make([]string, 0, len(extPkgs))
	for pkg := range extPkgs {
		pkgNames = append(pkgNames, pkg)
	}
	sort.Strings(pkgNames)
	for _, pkg := range pkgNames {
		funcs := make([]string, 0, len(extPkgs[pkg]))
		for fn := range extPkgs[pkg] {
			funcs = append(funcs, snakeCase(fn))
		}
		sort.Strings(funcs)
		b.WriteString(fmt.Sprintf("from app.services.%s import %s\n", pkg, strings.Join(funcs, ", ")))
	}
}
