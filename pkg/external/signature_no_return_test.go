//ff:func feature=external type=test control=sequence
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증
package external

import (
	"testing"
)

func TestSignatureNoReturn(t *testing.T) {
	m := methodInfo{Name: "DeleteItem", Params: []paramInfo{{Name: "id", GoType: "int64", In: "path"}}}
	got := m.signature()
	want := "DeleteItem(ctx context.Context, id int64) error"
	if got != want {
		t.Errorf("signature = %q, want %q", got, want)
	}
}
