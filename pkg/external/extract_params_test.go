//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증

package external

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func strSchema() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
}

func intSchema() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int64"}}
}

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

func TestExtractPathParamsNone(t *testing.T) {
	op := &openapi3.Operation{}
	if got := extractPathParams(op); got != nil {
		t.Errorf("expected nil, got %+v", got)
	}
}

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

func TestExtractBodyParamsNoBody(t *testing.T) {
	if got := extractBodyParams(&openapi3.Operation{}); got != nil {
		t.Errorf("expected nil for missing request body, got %+v", got)
	}
}

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

func TestDetectReturnTypeNo200(t *testing.T) {
	op := &openapi3.Operation{OperationID: "x", Responses: openapi3.NewResponses()}
	if got := detectReturnType(op); got != "" {
		t.Errorf("expected empty return type, got %q", got)
	}
}

func TestSortedKeys(t *testing.T) {
	m := openapi3.Schemas{"c": strSchema(), "a": strSchema(), "b": strSchema()}
	got := sortedKeys(m)
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortedKeys = %v, want %v", got, want)
	}
}

func TestSortedPathKeys(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/zebra", &openapi3.PathItem{})
	doc.Paths.Set("/apple", &openapi3.PathItem{})
	doc.Paths.Set("/mango", &openapi3.PathItem{})
	got := sortedPathKeys(doc)
	want := []string{"/apple", "/mango", "/zebra"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortedPathKeys = %v, want %v", got, want)
	}
}

func TestFindOperation(t *testing.T) {
	getOp := &openapi3.Operation{OperationID: "list"}
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/items", &openapi3.PathItem{Get: getOp})

	if got := findOperation(doc, "GET", "/items"); got != getOp {
		t.Errorf("findOperation(GET /items) = %v, want %v", got, getOp)
	}
	if got := findOperation(doc, "POST", "/items"); got != nil {
		t.Errorf("findOperation(POST /items) = %v, want nil", got)
	}
	if got := findOperation(doc, "GET", "/missing"); got != nil {
		t.Errorf("findOperation(GET /missing) = %v, want nil", got)
	}
}

func TestFindOperationNilPaths(t *testing.T) {
	if got := findOperation(&openapi3.T{}, "GET", "/x"); got != nil {
		t.Errorf("expected nil for doc with nil Paths, got %v", got)
	}
}
