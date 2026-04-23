//ff:func feature=funcspec type=parser control=sequence
//ff:what loadTypeMapForDir — 캐시 미스 시 디렉토리 타입을 수집하고 diag 를 seenDir 기준으로 누적

package funcspec

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// loadTypeMapForDir returns the cached type map for dir, collecting it on
// first access. Newly produced diagnostics are appended to diags only once
// per directory (tracked in seenDir).
func loadTypeMapForDir(dir string, cache map[string]map[string][]Field, seenDir map[string]struct{}, diags []diagnostic.Diagnostic) (map[string][]Field, []diagnostic.Diagnostic) {
	if typeMap, ok := cache[dir]; ok {
		return typeMap, diags
	}
	typeMap, d := collectPackageTypes(dir)
	cache[dir] = typeMap
	if _, seen := seenDir[dir]; !seen {
		diags = append(diags, d...)
		seenDir[dir] = struct{}{}
	}
	return typeMap, diags
}
