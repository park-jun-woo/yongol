//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractBodyParams(t *testing.T) {
	body := openapi3.NewRequestBody().WithJSONSchema(&openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"name":    strSchema(),
			"user_id": intSchema(),
		},
	})
	op := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: body}}

	got := extractBodyParams(op)
	// sortedKeys orders properties alphabetically: name, user_id.
	want := []paramInfo{
		{Name: "name", GoType: "string", In: "body"},
		{Name: "userID", GoType: "int64", In: "body"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractBodyParams = %+v, want %+v", got, want)
	}
}
