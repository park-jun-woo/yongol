//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectOpIDs — OpenAPI operationId 집합 수집 분기(nil doc/paths, 빈 ID 스킵, 다중 verb) 검증
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectOpIDs(t *testing.T) {
	// nil doc → empty set.
	if got := collectOpIDs(nil); len(got) != 0 {
		t.Fatalf("nil doc: expected empty, got %d", len(got))
	}
	// nil paths → empty set.
	if got := collectOpIDs(&openapi3.T{}); len(got) != 0 {
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
	// Empty operationId (skipped) plus an undefined verb (nil op, skipped).
	paths.Set("/skip", &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: ""},
	})

	out := collectOpIDs(&openapi3.T{Paths: paths})

	want := []string{"ListItems", "CreateItem", "ReplaceItem", "DeleteItem", "PatchItem"}
	if len(out) != len(want) {
		t.Fatalf("expected %d ids, got %d: %+v", len(want), len(out), out)
	}
	for _, id := range want {
		if _, ok := out[id]; !ok {
			t.Errorf("missing operationId %q", id)
		}
	}
	if _, ok := out[""]; ok {
		t.Error("empty operationId should not be collected")
	}
}
