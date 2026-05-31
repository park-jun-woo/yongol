//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증
package external

import (
	"strings"
	"testing"
)

func TestImplementationPostWithBody(t *testing.T) {
	m := methodInfo{
		Name:       "CreateItem",
		HTTPMethod: "POST",
		Path:       "/items",
		Params:     []paramInfo{{Name: "name", GoType: "string", In: "body"}},
		ReturnType: "",
	}
	code := m.implementation("Client")
	for _, want := range []string{
		"func (c *Client) CreateItem(ctx context.Context, name string) error {",
		`body := map[string]any{"name": name}`,
		`return c.do(ctx, "POST", "/items", body, nil)`,
	} {
		if !strings.Contains(code, want) {
			t.Errorf("implementation missing %q\n got:\n%s", want, code)
		}
	}
}
