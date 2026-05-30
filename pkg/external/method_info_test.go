//ff:func feature=external type=test control=sequence
//ff:what TestMethodInfo* — pathParams/bodyParams/buildPathExpr/signature/implementation 검증

package external

import (
	"strings"
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

func TestBuildPathExprNoParams(t *testing.T) {
	m := methodInfo{Path: "/items"}
	if got := m.buildPathExpr(); got != `"/items"` {
		t.Errorf("buildPathExpr = %q, want \"/items\"", got)
	}
}

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

func TestSignatureNoReturn(t *testing.T) {
	m := methodInfo{Name: "DeleteItem", Params: []paramInfo{{Name: "id", GoType: "int64", In: "path"}}}
	got := m.signature()
	want := "DeleteItem(ctx context.Context, id int64) error"
	if got != want {
		t.Errorf("signature = %q, want %q", got, want)
	}
}

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
