//ff:func feature=external type=test control=sequence
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증
package external

import (
	"testing"
)

func TestMethodInfoPathAndBodyParams(t *testing.T) {
	m := methodInfo{Params: []paramInfo{
		{Name: "id", In: "path"},
		{Name: "name", In: "body"},
		{Name: "age", In: "body"},
	}}
	pp := m.pathParams()
	if len(pp) != 1 || pp[0].Name != "id" {
		t.Errorf("pathParams = %+v", pp)
	}
	bp := m.bodyParams()
	if len(bp) != 2 || bp[0].Name != "name" || bp[1].Name != "age" {
		t.Errorf("bodyParams = %+v", bp)
	}
}
