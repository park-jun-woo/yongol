//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.applyOperation — 매칭된 verb/op 에서 Method, success status, params, body/resp 필드를 methodGen 에 적재

package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"

	yopenapi "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func (g *methodGen) applyOperation(pathItem *openapi3.PathItem, v verbOp, operationId string) {
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
	g.addAllParams(pathItem, v, operationId)
	g.extractBodyFormats(v.op)
	g.extractRespFields(v.op)
}
