//ff:func feature=validate type=util control=iteration dimension=1 topic=func-check
//ff:what findFuncSpec — callKey(pkg.camel) 로 ProjectFuncSpecs → YongolPkgSpecs 순으로 조회

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// findFuncSpec returns the first FuncSpec whose key (pkg.camelName) matches
// callKey (already camelCase-bridged). ProjectFuncSpecs override
// YongolPkgSpecs. Case must be exact — SSaC callers are expected to build
// callKey via callFuncName + toCamelKey to bridge PascalCase @call ↔
// camelCase @func annotation.
func findFuncSpec(callKey string, project, yongolPkg []funcspec.FuncSpec) *funcspec.FuncSpec {
	for i := range project {
		k := project[i].Package + "." + project[i].Name
		if k == callKey {
			return &project[i]
		}
	}
	for i := range yongolPkg {
		k := yongolPkg[i].Package + "." + yongolPkg[i].Name
		if k == callKey {
			return &yongolPkg[i]
		}
	}
	return nil
}
