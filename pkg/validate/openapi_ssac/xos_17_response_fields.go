//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-17 — SSaC @response 필드가 OpenAPI response 스키마에 포함되는지 검증

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos17ResponseFields validates XOS-17: explicit @response field keys match an
// OpenAPI 2xx response schema property.
func xos17ResponseFields(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
		opProps := toSet(g.Schemas["OpenAPI.response."+fn.Name])
		if len(opProps) == 0 {
			continue
		}
		diags = append(diags, collectMissingProps(fn, keys, opProps, "XOS-17")...)
	}
	return diags
}
