//ff:func feature=external type=test control=sequence
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증
package external

import (
	"testing"
)

func TestSignatureWithReturn(t *testing.T) {
	m := methodInfo{
		Name:       "GetItem",
		Params:     []paramInfo{{Name: "id", GoType: "int64", In: "path"}},
		ReturnType: "GetItemResponse",
	}
	got := m.signature()
	want := "GetItem(ctx context.Context, id int64) (*GetItemResponse, error)"
	if got != want {
		t.Errorf("signature = %q, want %q", got, want)
	}
}
