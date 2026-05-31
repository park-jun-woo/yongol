//ff:func feature=external type=test control=sequence
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증
package external

import (
	"testing"
)

func TestBuildPathExprWithParams(t *testing.T) {
	m := methodInfo{
		Path:   "/items/{id}/sub/{sub}",
		Params: []paramInfo{{Name: "id", In: "path"}, {Name: "sub", In: "path"}},
	}
	got := m.buildPathExpr()
	want := `fmt.Sprintf("/items/%v/sub/%v", id, sub)`
	if got != want {
		t.Errorf("buildPathExpr = %q, want %q", got, want)
	}
}
