//ff:func feature=validate type=rule control=iteration dimension=3 topic=openapi-ssac
//ff:what XOS-70 — @response integer field (literal or variable binding, required or optional) must have format: int64

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos70ResponseLiteralIntFormat validates that when a @response field maps to an
// OpenAPI integer property, that property declares format: int64.
//
// codegen binds integer response fields to int64: integer literals become
// ptrOf(int64(<lit>)) (optional) or int64(<lit>) (required); variable bindings
// become &<var> / <var> whose Go type is int64 (DDL columns, COUNT/Func results).
// oapi-codegen generates int/*int when format is absent and int64/*int64 when
// format: int64. Without int64 the generated types are incompatible and the
// build fails. DDL-backed integer fields are already forced to format: int64 by
// XDO-77; this rule additionally covers non-DDL (Func/COUNT) integer responses.
func xos70ResponseLiteralIntFormat(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.ResponseConstraints == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		rc := fs.ResponseConstraints[fn.Name]
		if rc == nil {
			continue
		}
		for _, seq := range fn.Sequences {
			if seq.Type != "response" {
				continue
			}
			for key, value := range seq.Fields {
				if diag, ok := xos70FieldDiag(fn, seq.Line, key, value, rc); ok {
					diags = append(diags, diag)
				}
			}
		}
	}
	return diags
}
