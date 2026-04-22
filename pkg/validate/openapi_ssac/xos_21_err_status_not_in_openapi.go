//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-21 — SSaC @empty/@exists/@state/@auth/@call ErrStatus가 OpenAPI 응답에 정의되어 있는지 검증

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos21ErrStatusNotInOpenAPI validates XOS-21: every ErrStatus used by a
// guard-style sequence (@empty/@exists/@state/@auth) is defined as a response
// on the matching OpenAPI operation.
func xos21ErrStatusNotInOpenAPI(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil {
		return nil
	}
	opMap := buildOperationMap(fs.OpenAPIDoc)
	allFuncSpecs := append(fs.ProjectFuncSpecs, fs.YongolPkgSpecs...)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		op := opMap[fn.Name]
		if op == nil || op.Responses == nil {
			continue
		}
		diags = append(diags, checkFuncErrStatus(fn.FileName, fn.Name, fn.Sequences, op, allFuncSpecs)...)
	}
	return diags
}
