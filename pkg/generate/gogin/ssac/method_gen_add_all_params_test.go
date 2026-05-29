//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.addAllParams 단위 테스트 (path-level + operation-level 파라미터 모두 등록)

package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestMethodGenAddAllParams(t *testing.T) {
	op := &openapi3.Operation{
		OperationID: "GetThing",
		Parameters: openapi3.Parameters{
			paramRef("cursor", "query", false, "string"),
		},
	}
	pathItem := &openapi3.PathItem{
		Parameters: openapi3.Parameters{
			paramRef("id", "path", true, "integer"),
		},
		Get: op,
	}
	g := newParamGen()
	g.addAllParams(pathItem, verbOp{method: "GET", op: op}, "GetThing")

	if !g.PathParams["id"] {
		t.Errorf("path-level param id missing: %v", g.PathParams)
	}
	if _, ok := g.QueryParams["cursor"]; !ok {
		t.Errorf("operation-level query param cursor missing: %v", g.QueryParams)
	}
}
