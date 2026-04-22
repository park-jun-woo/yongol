//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what appendFuncSpecEntries — FuncSpec 목록을 "pkg.Name" 키로 calls 맵에 추가

package contract

import "github.com/park-jun-woo/yongol/pkg/parser/funcspec"

// appendFuncSpecEntries adds "<Package>.<Name>" keys to calls for every
// FuncSpec whose package and name are both non-empty. Existing entries
// are overwritten harmlessly because the map value is a presence flag.
func appendFuncSpecEntries(specs []funcspec.FuncSpec, calls map[string]bool) {
	for _, s := range specs {
		if s.Package != "" && s.Name != "" {
			calls[s.Package+"."+s.Name] = true
		}
	}
}
