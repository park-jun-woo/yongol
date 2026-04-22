//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what lookupFuncSpecErrStatus — @call model 에 대응하는 FuncSpec @error status 찾기

package openapi_ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// lookupFuncSpecErrStatus finds the @error status for a @call model (e.g. "billing.Spend").
func lookupFuncSpecErrStatus(model string, specs []funcspec.FuncSpec) int {
	parts := strings.SplitN(model, ".", 2)
	if len(parts) < 2 {
		return 0
	}
	pkg, fn := parts[0], parts[1]
	fnLower := strings.ToLower(fn[:1]) + fn[1:] // Spend → spend
	for _, sp := range specs {
		if sp.Package == pkg && (sp.Name == fn || sp.Name == fnLower) && sp.ErrStatus != 0 {
			return sp.ErrStatus
		}
	}
	return 0
}
