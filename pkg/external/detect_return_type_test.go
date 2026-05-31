//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestDetectReturnType(t *testing.T) {
	resp := openapi3.NewResponse().WithJSONSchema(&openapi3.Schema{Type: &openapi3.Types{"object"}})
	op := &openapi3.Operation{
		OperationID: "get_user",
		Responses:   openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{Value: resp})),
	}
	if got := detectReturnType(op); got != "GetUserResponse" {
		t.Errorf("detectReturnType = %q, want GetUserResponse", got)
	}
}
