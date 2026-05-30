//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestBuildOperationMethodMap — operationId → (method, op) 맵 생성 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildOperationMethodMap(t *testing.T) {
	// nil doc → empty map.
	if got := buildOperationMethodMap(nil); len(got) != 0 {
		t.Fatalf("nil doc: expected empty, got %d", len(got))
	}
	// nil paths → empty map.
	if got := buildOperationMethodMap(&openapi3.T{}); len(got) != 0 {
		t.Fatalf("nil paths: expected empty, got %d", len(got))
	}

	paths := openapi3.NewPaths()
	paths.Set("/items", &openapi3.PathItem{
		Get:    &openapi3.Operation{OperationID: "ListItems"},
		Post:   &openapi3.Operation{OperationID: "CreateItem"},
		Put:    &openapi3.Operation{OperationID: "ReplaceItem"},
		Delete: &openapi3.Operation{OperationID: "DeleteItem"},
		Patch:  &openapi3.Operation{OperationID: "PatchItem"},
	})
	// Path with an operation lacking an OperationID (skipped) and a nil verb.
	paths.Set("/skip", &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: ""},
	})

	doc := &openapi3.T{Paths: paths}
	out := buildOperationMethodMap(doc)

	want := map[string]string{
		"ListItems":   "GET",
		"CreateItem":  "POST",
		"ReplaceItem": "PUT",
		"DeleteItem":  "DELETE",
		"PatchItem":   "PATCH",
	}
	if len(out) != len(want) {
		t.Fatalf("expected %d entries, got %d: %+v", len(want), len(out), out)
	}
	for id, method := range want {
		e, ok := out[id]
		if !ok {
			t.Fatalf("missing %q", id)
		}
		if e.method != method {
			t.Errorf("%s method = %q, want %q", id, e.method, method)
		}
		if e.op == nil {
			t.Errorf("%s op is nil", id)
		}
	}
}
