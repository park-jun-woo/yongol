//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what checkFuncErrStatus — verifies that ErrStatus codes in a single SSaC function are defined in OpenAPI

package openapi_ssac

import (
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkFuncErrStatus reports each guard-style sequence whose effective HTTP
// status code is not defined on the OpenAPI operation.
// For @call, the effective status is: seq.ErrStatus > FuncSpec @error > default 500.
func checkFuncErrStatus(file, funcName string, seqs []ssac.Sequence, op *openapi3.Operation, funcSpecs []funcspec.FuncSpec) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range seqs {
		status, ok := resolveSeqErrStatus(seq, funcSpecs)
		if !ok {
			continue
		}
		if op.Responses.Status(status) != nil {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    seq.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XOS-21] SSaC @" + seq.Type + " in " + funcName + " uses HTTP " + strconv.Itoa(status) + " but OpenAPI has no " + strconv.Itoa(status) + " response",
			Advice:  "Add a " + strconv.Itoa(status) + " response to the OpenAPI " + funcName + " responses",
		})
	}
	return diags
}
