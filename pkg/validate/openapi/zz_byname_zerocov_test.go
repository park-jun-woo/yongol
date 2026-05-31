//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what TestByName_ZeroCov — O-6 스키마 워커들을 이름으로 직접 호출해 커버리지 귀속

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// byNameO06Doc builds an OpenAPI doc with component schemas, a request body,
// responses, array items and nested children — enough to exercise every o06
// walker by name.
func byNameO06Doc() *openapi3.T {
	itemSchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: openapi3.Schemas{"id": {Value: openapi3.NewSchema()}},
		Required:   []string{"phantom"},
	}}
	arraySchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: itemSchema,
	}}
	objSchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"name":  {Value: openapi3.NewSchema()},
			"items": arraySchema,
		},
		Required: []string{"name", "missing"},
	}}

	reqBody := &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
		WithJSONSchemaRef(objSchema)}

	resp := openapi3.NewResponse().WithJSONSchemaRef(arraySchema)
	resps := openapi3.NewResponses()
	resps.Set("200", &openapi3.ResponseRef{Value: resp})

	op := &openapi3.Operation{OperationID: "CreateItem", RequestBody: reqBody, Responses: resps}
	pi := &openapi3.PathItem{Post: op}

	return &openapi3.T{
		Components: &openapi3.Components{Schemas: openapi3.Schemas{"Workflow": objSchema}},
		Paths:      openapi3.NewPaths(openapi3.WithPath("/items", pi)),
	}
}

func byNameO06FS(doc *openapi3.T) *yongol.Fullstack {
	return &yongol.Fullstack{
		OpenAPIDoc: doc,
		OpenAPILines: &oapiparser.LineIndex{
			Schemas:          map[string]int{"Workflow": 10},
			SchemaProperties: map[string]map[string]int{"Workflow": {"name": 11}},
			Paths:            map[string]int{},
			Operations:       map[string]int{},
		},
	}
}

func TestByNameO06Collectors_ZeroCov(t *testing.T) {
	doc := byNameO06Doc()
	fs := byNameO06FS(doc)

	all := o06CollectAllSchemas(fs)
	if len(all) == 0 {
		t.Fatalf("o06CollectAllSchemas empty")
	}
	if o06CollectAllSchemas(nil) != nil {
		t.Errorf("o06CollectAllSchemas(nil) should be nil")
	}
	if o06CollectAllSchemas(&yongol.Fullstack{}) != nil {
		t.Errorf("o06CollectAllSchemas(no doc) should be nil")
	}

	visited := map[*openapi3.Schema]bool{}
	comp := o06CollectComponentSchemas(doc, visited, nil)
	if len(comp) == 0 {
		t.Errorf("o06CollectComponentSchemas empty")
	}
	visited2 := map[*openapi3.Schema]bool{}
	pathAcc := o06CollectPathSchemas(doc, visited2, nil)
	_ = pathAcc

	// item schemas walk for the one path item.
	for _, pi := range doc.Paths.Map() {
		visited3 := map[*openapi3.Schema]bool{}
		_ = o06CollectItemSchemas(pi, visited3, nil)
	}
}

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

func TestByNameO06CheckAndLine_ZeroCov(t *testing.T) {
	doc := byNameO06Doc()
	fs := byNameO06FS(doc)

	entry := o06SchemaEntry{
		schema:     doc.Components.Schemas["Workflow"].Value,
		schemaName: "Workflow",
	}
	diags := o06CheckSchemaRequired(fs, entry)
	// "missing" is required but not a property → a diagnostic is expected.
	if len(diags) == 0 {
		t.Errorf("o06CheckSchemaRequired expected diagnostics for dangling required")
	}

	if line := o06RequiredLine(fs, "Workflow", "name"); line == 0 {
		t.Errorf("o06RequiredLine fallback returned 0")
	}
	_ = o06RequiredLine(fs, "Unknown", "x")
}

func TestByNameXOE01_ZeroCov(t *testing.T) {
	doc := byNameO06Doc()
	fs := byNameO06FS(doc)
	_ = xoe01ErrorResponseRequired(fs)
	_ = xoe01ErrorResponseRequired(&yongol.Fullstack{})
}
