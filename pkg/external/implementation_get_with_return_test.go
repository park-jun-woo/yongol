//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증
package external

import (
	"strings"
	"testing"
)

func TestImplementationGetWithReturn(t *testing.T) {
	m := methodInfo{
		Name:       "GetItem",
		HTTPMethod: "GET",
		Path:       "/items/{id}",
		Params:     []paramInfo{{Name: "id", GoType: "int64", In: "path"}},
		ReturnType: "GetItemResponse",
	}
	code := m.implementation("Client")
	for _, want := range []string{
		"func (c *Client) GetItem(ctx context.Context, id int64) (*GetItemResponse, error) {",
		`fmt.Sprintf("/items/%v", id)`,
		"var resp GetItemResponse",
		`c.do(ctx, "GET",`,
	} {
		if !strings.Contains(code, want) {
			t.Errorf("implementation missing %q\n got:\n%s", want, code)
		}
	}
	// GET must not build a body map.
	if strings.Contains(code, "body := map[string]any{") {
		t.Error("GET implementation should not emit a body map")
	}
}
