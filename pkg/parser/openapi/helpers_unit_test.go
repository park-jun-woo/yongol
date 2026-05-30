//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

func opWith2xx(codes ...int) *openapi3.Operation {
	op := &openapi3.Operation{Responses: openapi3.NewResponses()}
	for _, c := range codes {
		op.Responses.Set(itoaCode(c), &openapi3.ResponseRef{Value: &openapi3.Response{}})
	}
	return op
}

func itoaCode(c int) string {
	// small local itoa for 3-digit status codes
	return string(rune('0'+c/100)) + string(rune('0'+(c/10)%10)) + string(rune('0'+c%10))
}

func TestDeclared2xx(t *testing.T) {
	op := opWith2xx(200, 201)
	op.Responses.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{}})
	op.Responses.Set("default", &openapi3.ResponseRef{Value: &openapi3.Response{}})
	set := declared2xx(op)
	if !set[200] || !set[201] {
		t.Errorf("missing 2xx codes: %v", set)
	}
	if set[404] {
		t.Error("404 should not be in 2xx set")
	}
	if len(set) != 2 {
		t.Errorf("expected 2 codes, got %v", set)
	}
	// nil op → empty map.
	if got := declared2xx(nil); len(got) != 0 {
		t.Errorf("nil op → %v", got)
	}
	// nil responses → empty map.
	if got := declared2xx(&openapi3.Operation{}); len(got) != 0 {
		t.Errorf("nil responses → %v", got)
	}
}

func TestDeclared2xxExported(t *testing.T) {
	set := Declared2xx(opWith2xx(200))
	if !set[200] {
		t.Errorf("Declared2xx missing 200: %v", set)
	}
}

func TestExtractSchemaConstraints(t *testing.T) {
	schema := &openapi3.Schema{
		Required: []string{"id"},
		Properties: openapi3.Schemas{
			"id":   {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"bad":  {Value: nil}, // skipped
		},
	}
	fields := extractSchemaConstraints(schema)
	if !fields["id"].Required {
		t.Error("id should be required")
	}
	if fields["name"].Required {
		t.Error("name should not be required")
	}
	if fields["id"].Type != "integer" {
		t.Errorf("id type = %q", fields["id"].Type)
	}
	if _, ok := fields["bad"]; ok {
		t.Error("nil-value property should be skipped")
	}
}

const respYAML = `
content:
  application/json:
    schema:
      properties:
        id:
          type: integer
        name:
          type: string
`

func parseNode(t *testing.T, src string) *yaml.Node {
	t.Helper()
	var root yaml.Node
	if err := yaml.Unmarshal([]byte(src), &root); err != nil {
		t.Fatal(err)
	}
	return root.Content[0]
}

func TestSchemaPropsOfBody(t *testing.T) {
	body := parseNode(t, respYAML)
	props := schemaPropsOfBody(body)
	if props == nil || props.Kind != yaml.MappingNode {
		t.Fatalf("expected properties mapping, got %v", props)
	}
	lines := collectPropertyLines(props)
	if _, ok := lines["id"]; !ok {
		t.Errorf("missing id in property lines: %v", lines)
	}

	// Missing content → nil.
	noContent := parseNode(t, "summary: hello\n")
	if got := schemaPropsOfBody(noContent); got != nil {
		t.Errorf("no content → %v, want nil", got)
	}
}

func TestIndexFirst2xxResponse(t *testing.T) {
	respsYAML := "" +
		"\"200\":\n" +
		"  content:\n" +
		"    application/json:\n" +
		"      schema:\n" +
		"        properties:\n" +
		"          id:\n" +
		"            type: integer\n" +
		"\"404\":\n" +
		"  description: not found\n"
	resps := parseNode(t, respsYAML)
	idx := &LineIndex{ResponseFields: map[string]map[string]int{}}
	indexFirst2xxResponse(resps, "GetThing", idx)
	if _, ok := idx.ResponseFields["GetThing"]["id"]; !ok {
		t.Errorf("expected id field line indexed, got %v", idx.ResponseFields["GetThing"])
	}
}
