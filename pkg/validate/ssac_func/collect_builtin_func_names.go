//ff:func feature=validate type=util control=iteration dimension=1 topic=func-check
//ff:what collectBuiltinFuncNames — builtin 패키지의 함수 이름을 YongolPkgSpecs + Ground에서 병합 수집

package ssac_func

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectBuiltinFuncNames merges function names from YongolPkgSpecs and
// Ground Func.spec for a given builtin package. Returns sorted, deduplicated
// PascalCase names.
func collectBuiltinFuncNames(pkg string, fs *yongol.Fullstack) []string {
	seen := map[string]bool{}
	for _, name := range builtinFuncNames(pkg, fs.YongolPkgSpecs) {
		seen[name] = true
	}
	collectGroundFuncNames(pkg, fs, seen)
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
