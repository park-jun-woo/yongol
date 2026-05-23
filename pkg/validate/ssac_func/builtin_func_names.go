//ff:func feature=validate type=util control=iteration dimension=1 topic=func-check
//ff:what builtinFuncNames — 내장 패키지의 SSaC @call 가능 함수 PascalCase 이름 목록 반환

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// builtinFuncNames returns the PascalCase names of functions available in the
// given builtin package, extracted from YongolPkgSpecs. The returned list
// preserves the order they appear in specs. FuncSpec.Name is camelCase; this
// function upper-cases the first letter to match SSaC @call convention.
func builtinFuncNames(pkg string, specs []funcspec.FuncSpec) []string {
	var names []string
	for _, sp := range specs {
		if sp.Package != pkg {
			continue
		}
		names = append(names, ucFirst(sp.Name))
	}
	return names
}

// ucFirst upper-cases the first byte of s. ASCII-only (sufficient for Go
// identifier names).
func ucFirst(s string) string {
	if s == "" {
		return s
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-32) + s[1:]
	}
	return s
}
