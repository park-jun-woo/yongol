//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-22 — verifies that functions with @response have a 2xx response code in OpenAPI

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos22ResponseNo2xx validates XOS-22: a SSaC func with @response must have at
// least one explicit 2xx response code defined in OpenAPI.
func xos22ResponseNo2xx(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	opMap := buildOperationMapAll(fs)
	if len(opMap) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if !hasResponseSequence(fn) {
			continue
		}
		op := opMap[fn.Name]
		if op == nil {
			continue
		}
		if hasExplicit2xx(op) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        fn.FileName,
			Line:        fn.Line,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     "[XOS-22] SSaC " + fn.Name + " has @response but OpenAPI defines no explicit 2xx response",
			Advice:      "Declare a 2xx response (200, 201, 204, etc.) in the OpenAPI " + fn.Name + " responses",
			OperationID: fn.Name,
		})
	}
	return diags
}
