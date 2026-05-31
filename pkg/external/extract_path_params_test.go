//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractPathParams(t *testing.T) {
	op := &openapi3.Operation{
		Parameters: openapi3.Parameters{
			{Value: &openapi3.Parameter{Name: "user_id", In: "path", Schema: intSchema()}},
			{Value: &openapi3.Parameter{Name: "q", In: "query", Schema: strSchema()}},
		},
	}
	got := extractPathParams(op)
	want := []paramInfo{{Name: "userID", GoType: "int64", In: "path"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractPathParams = %+v, want %+v", got, want)
	}
}
