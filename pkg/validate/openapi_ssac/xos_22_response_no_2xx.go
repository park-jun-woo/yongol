//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-22 — @response 있는 함수에 OpenAPI 2xx 응답 코드 존재 여부 검증

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos22ResponseNo2xx validates XOS-22: a SSaC func with @response must have at
// least one explicit 2xx response code defined in OpenAPI.
func xos22ResponseNo2xx(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil {
		return nil
	}
	opMap := buildOperationMap(fs.OpenAPIDoc)
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
			File:    fn.FileName,
			Line:    fn.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XOS-22] SSaC " + fn.Name + " has @response but OpenAPI defines no explicit 2xx response",
			Advice:  "OpenAPI " + fn.Name + " responses 에 2xx 응답(200/201/204 등)을 명시하세요",
		})
	}
	return diags
}
