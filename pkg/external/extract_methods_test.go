//ff:func feature=external type=test control=sequence
//ff:what TestExtractMethods/buildMethodInfo/extractResponseTypes — 메서드·응답타입 추출 검증

package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func sampleDoc() *openapi3.T {
	getResp := openapi3.NewResponse().WithJSONSchema(&openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id":   intSchema(),
			"name": strSchema(),
		},
	})
	getOp := &openapi3.Operation{
		OperationID: "get_item",
		Parameters: openapi3.Parameters{
			{Value: &openapi3.Parameter{Name: "item_id", In: "path", Schema: intSchema()}},
		},
		Responses: openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{Value: getResp})),
	}
	postOp := &openapi3.Operation{
		OperationID: "create_item",
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchema(&openapi3.Schema{
			Type:       &openapi3.Types{"object"},
			Properties: openapi3.Schemas{"name": strSchema()},
		})},
		Responses: openapi3.NewResponses(), // no 200 => error-only
	}

	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/items/{item_id}", &openapi3.PathItem{Get: getOp})
	doc.Paths.Set("/items", &openapi3.PathItem{Post: postOp})
	return doc
}

func TestBuildMethodInfo(t *testing.T) {
	doc := sampleDoc()
	getOp := doc.Paths.Map()["/items/{item_id}"].Get

	m := buildMethodInfo(getOp, "GET", "/items/{item_id}")
	if m.Name != "GetItem" {
		t.Errorf("Name = %q, want GetItem", m.Name)
	}
	if m.HTTPMethod != "GET" || m.Path != "/items/{item_id}" {
		t.Errorf("method/path mismatch: %+v", m)
	}
	if m.ReturnType != "GetItemResponse" {
		t.Errorf("ReturnType = %q, want GetItemResponse", m.ReturnType)
	}
	if len(m.Params) != 1 || m.Params[0].Name != "itemID" {
		t.Errorf("params = %+v", m.Params)
	}
}

func TestExtractMethods(t *testing.T) {
	doc := sampleDoc()
	methods := extractMethods(doc)
	// sorted by path: /items (POST create_item), /items/{item_id} (GET get_item)
	if len(methods) != 2 {
		t.Fatalf("expected 2 methods, got %d: %+v", len(methods), methods)
	}
	if methods[0].Name != "CreateItem" || methods[0].HTTPMethod != "POST" {
		t.Errorf("methods[0] = %+v, want CreateItem/POST", methods[0])
	}
	if methods[1].Name != "GetItem" || methods[1].HTTPMethod != "GET" {
		t.Errorf("methods[1] = %+v, want GetItem/GET", methods[1])
	}
}

func TestExtractMethodsNilPaths(t *testing.T) {
	if got := extractMethods(&openapi3.T{}); len(got) != 0 {
		t.Errorf("expected no methods for nil Paths, got %+v", got)
	}
}

func TestExtractResponseTypes(t *testing.T) {
	doc := sampleDoc()
	methods := extractMethods(doc)
	types := extractResponseTypes("svc", methods, doc)

	// Only get_item has a 200 JSON object response with properties.
	if len(types) != 1 {
		t.Fatalf("expected 1 response type, got %d: %+v", len(types), types)
	}
	st := types[0]
	if st.Name != "GetItemResponse" {
		t.Errorf("Name = %q, want GetItemResponse", st.Name)
	}
	if len(st.Fields) != 2 {
		t.Fatalf("expected 2 fields, got %+v", st.Fields)
	}
	// sortedKeys => id, name
	if st.Fields[0].Name != "ID" || st.Fields[0].JSONName != "id" || st.Fields[0].GoType != "int64" {
		t.Errorf("fields[0] = %+v", st.Fields[0])
	}
	if st.Fields[1].Name != "Name" || st.Fields[1].JSONName != "name" {
		t.Errorf("fields[1] = %+v", st.Fields[1])
	}
}
