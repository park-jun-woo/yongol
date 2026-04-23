//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-82 — OpenAPI 에 선언된 2xx 중 yongol 이 emit 하지 않는 것이 있음

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos82UnreachableSuccessStatus validates XOS-82: the OpenAPI operation
// declares multiple 2xx responses but yongol emits only the one selected
// by DeriveSuccessStatus. The remaining 2xx codes are unreachable from
// the generated handler. Authors may have pre-declared codes for
// forward-compat (e.g. future `@upsert` that returns either 200 or 201);
// WARNING keeps the operation visible without blocking codegen. The
// warning also nudges authors to trim genuinely dead declarations.
func xos82UnreachableSuccessStatus(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil {
		return nil
	}
	opMap := buildOperationMethodMap(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, xos82CheckFunc(fn, opMap)...)
	}
	return diags
}
