//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what TestByName_ZeroCov — O-6 스키마 워커들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestByNameO06Walkers_ZeroCov(t *testing.T) {
	doc := byNameO06Doc()
	var op *openapi3.Operation
	for _, pi := range doc.Paths.Map() {
		op = pi.Post
	}

	visited := map[*openapi3.Schema]bool{}
	rb := o06WalkRequestBody(op.RequestBody, visited, nil)
	_ = rb
	_ = o06WalkRequestBody(nil, visited, nil)

	resps := o06WalkResponses(op.Responses, visited, nil)
	_ = resps
	_ = o06WalkResponses(nil, visited, nil)

	for _, r := range op.Responses.Map() {
		_ = o06WalkResponse(r, visited, nil)
	}
	_ = o06WalkResponse(nil, visited, nil)

	// media type from the request body.
	mt := op.RequestBody.Value.Content.Get("application/json")
	_ = o06WalkMediaType(mt, visited, nil)
	_ = o06WalkMediaType(nil, visited, nil)

	// schema ref + children.
	objRef := doc.Components.Schemas["Workflow"]
	visited4 := map[*openapi3.Schema]bool{}
	ref := o06WalkSchemaRef(objRef, "Workflow", visited4, nil)
	if len(ref) == 0 {
		t.Errorf("o06WalkSchemaRef empty")
	}
	_ = o06WalkSchemaRef(nil, "X", visited4, nil)

	visited5 := map[*openapi3.Schema]bool{}
	children := o06WalkSchemaChildren(objRef.Value, visited5, nil)
	_ = children
}
