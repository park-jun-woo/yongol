//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestDetectReturnTypeNoJSONContent(t *testing.T) {
	// 200 present but with no application/json content -> empty return type.
	op := &openapi3.Operation{
		OperationID: "x",
		Responses: openapi3.NewResponses(openapi3.WithStatus(200,
			&openapi3.ResponseRef{Value: openapi3.NewResponse()})),
	}
	if got := detectReturnType(op); got != "" {
		t.Errorf("expected empty return type for no JSON content, got %q", got)
	}
}
