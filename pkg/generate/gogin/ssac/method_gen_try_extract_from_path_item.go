//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.tryExtractFromPathItem — PathItem의 verb 중 매칭되는 op 찾아 메타데이터 적재

package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"

	yopenapi "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// tryExtractFromPathItem scans each verb on pathItem. When it finds the
// operation whose OperationID matches, it populates g's metadata maps from
// that op plus the path-level parameters and returns true. Returns false
// when no verb matches — extractFromOpenAPI keeps scanning.
func (g *methodGen) tryExtractFromPathItem(pathItem *openapi3.PathItem, operationId string) bool {
	type verbOp struct {
		method string
		op     *openapi3.Operation
	}
	verbs := []verbOp{
		{"GET", pathItem.Get},
		{"POST", pathItem.Post},
		{"PUT", pathItem.Put},
		{"DELETE", pathItem.Delete},
		{"PATCH", pathItem.Patch},
	}
	for _, v := range verbs {
		if v.op == nil || v.op.OperationID != operationId {
			continue
		}
		g.Method = v.method
		// Success status follows HTTP-method convention (POST→201, PUT→200,
		// DELETE→204, …) with a fallback to the lowest declared 2xx. This
		// replaces the prior hardcoded 200 in build_response /
		// build_field_response (BUG-004). If DeriveSuccessStatus returns
		// 0 (no 2xx declared), the default 200 set in newMethodGen
		// stays — the XOS-81 validator flags that case separately.
		if s := yopenapi.DeriveSuccessStatus(v.op, v.method); s != 0 {
			g.SuccessStatus = s
		}
		// path-level + operation-level parameters
		for _, p := range pathItem.Parameters {
			g.addParam(p, operationId)
		}
		for _, p := range v.op.Parameters {
			g.addParam(p, operationId)
		}
		// request body schema (for format-based casts)
		g.extractBodyFormats(v.op)
		// 200 response schema
		g.extractRespFields(v.op)
		return true
	}
	return false
}
