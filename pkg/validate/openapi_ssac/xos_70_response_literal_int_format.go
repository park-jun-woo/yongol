//ff:func feature=validate type=rule control=iteration dimension=3 topic=openapi-ssac
//ff:what XOS-70 — @response integer literal mapped to optional integer field must have format: int64

package openapi_ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos70ResponseLiteralIntFormat validates that when a @response field maps an
// integer literal to an optional integer property, the OpenAPI response schema
// declares format: int64.
//
// codegen emits ptrOf(int64(<lit>)) for integer literals. oapi-codegen generates
// *int when format is absent and *int64 when format: int64. Without int64 the
// generated types are incompatible and the build fails.
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
				if inferLiteral(value) != "int64" {
					continue
				}
				fc, ok := rc[key]
				if !ok {
					continue
				}
				if fc.Type != "integer" {
					continue
				}
				if fc.Required {
					continue
				}
				if fc.Format == "int64" {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  fn.FileName,
					Line:  seq.Line,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[XOS-70] %s @response field %q maps integer literal to optional integer without format: int64",
						fn.Name, key),
					Advice:      "Add format: int64 to OpenAPI response schema, or include the field in required",
					OperationID: fn.Name,
				})
			}
		}
	}
	return diags
}
