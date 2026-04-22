//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XSO-18 — OpenAPI response 필드가 SSaC 명시 @response에서 사용되는지 검증

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xso18ResponseFieldUsed validates XSO-18: every OpenAPI response property is
// referenced by the SSaC explicit @response field list for the same function.
func xso18ResponseFieldUsed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		keys := extractResponseFieldKeys(fn)
		if keys == nil {
			continue
		}
		opProps := g.Schemas["OpenAPI.response."+fn.Name]
		if len(opProps) == 0 {
			continue
		}
		used := toSet(keys)
		diags = append(diags, collectUnusedProps(fn, opProps, used, "XSO-18")...)
	}
	return diags
}
