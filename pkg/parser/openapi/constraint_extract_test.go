//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func arraySchemaRef(itemProps ...string) *openapi3.SchemaRef {
	props := openapi3.Schemas{}
	for _, p := range itemProps {
		props[p] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Properties: props}},
	}}
}

func jsonBodyOp(opID string, props openapi3.Schemas) *openapi3.Operation {
	schema := &openapi3.Schema{Properties: props}
	rb := openapi3.NewRequestBody().WithJSONSchema(schema)
	op := openapi3.NewOperation()
	op.OperationID = opID
	op.RequestBody = &openapi3.RequestBodyRef{Value: rb}
	return op
}

func jsonRespOp(opID string, props openapi3.Schemas) *openapi3.Operation {
	op := openapi3.NewOperation()
	op.OperationID = opID
	resp := openapi3.NewResponse().WithJSONSchema(&openapi3.Schema{Properties: props})
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{Value: resp})
	op.Responses = responses
	return op
}

func strProps(names ...string) openapi3.Schemas {
	s := openapi3.Schemas{}
	for _, n := range names {
		s[n] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	return s
}

func TestCollectItemFields(t *testing.T) {
	got := collectItemFields(arraySchemaRef("id", "name"))
	if !got["id"] || !got["name"] {
		t.Errorf("collectItemFields = %v", got)
	}
	// non-array yields nil
	scalar := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	if collectItemFields(scalar) != nil {
		t.Errorf("scalar should yield nil")
	}
	// array with nil items yields nil
	noItems := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"array"}}}
	if collectItemFields(noItems) != nil {
		t.Errorf("array without items should yield nil")
	}
}

func TestExtractArrayItemFields(t *testing.T) {
	schema := &openapi3.Schema{Properties: openapi3.Schemas{
		"items":  arraySchemaRef("id", "title"),
		"single": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}}
	got := extractArrayItemFields(schema)
	if len(got) != 1 {
		t.Fatalf("expected 1 array field, got %v", got)
	}
	if !got["items"]["id"] || !got["items"]["title"] {
		t.Errorf("items fields = %v", got["items"])
	}
}

func TestExtractBodyConstraints(t *testing.T) {
	op := jsonBodyOp("x", strProps("email", "name"))
	fc := extractBodyConstraints(op.RequestBody, "x")
	if len(fc) != 2 || fc["email"].Type != "string" {
		t.Errorf("constraints = %v", fc)
	}
	// nil content body
	empty := &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody()}
	if extractBodyConstraints(empty, "x") != nil {
		t.Errorf("empty body should yield nil")
	}
}

func TestExtractResponseFields(t *testing.T) {
	op := jsonRespOp("x", strProps("token"))
	fc := extractResponseFields(op)
	if len(fc) != 1 || fc["token"].Type != "string" {
		t.Errorf("response fields = %v", fc)
	}
	// no 2xx
	op2 := openapi3.NewOperation()
	op2.OperationID = "y"
	r := openapi3.NewResponses()
	r.Set("404", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	op2.Responses = r
	if extractResponseFields(op2) != nil {
		t.Errorf("no 2xx should yield nil")
	}
}

func TestCollectRequestConstraintsForOp(t *testing.T) {
	result := map[string]map[string]FieldConstraint{}
	op := jsonBodyOp("CreateUser", strProps("email"))
	collectRequestConstraintsForOp(result, op, nil)
	if result["CreateUser"]["email"].Type != "string" {
		t.Errorf("result = %v", result)
	}
	// op without operationId is skipped
	noID := jsonBodyOp("", strProps("x"))
	collectRequestConstraintsForOp(result, noID, nil)
	if _, ok := result[""]; ok {
		t.Errorf("empty opID should be skipped")
	}
}

func TestCollectResponseConstraintsForOp(t *testing.T) {
	result := map[string]map[string]FieldConstraint{}
	op := jsonRespOp("GetUser", strProps("id"))
	collectResponseConstraintsForOp(result, op, nil)
	if result["GetUser"]["id"].Type != "string" {
		t.Errorf("result = %v", result)
	}
}

func TestFillRequestAndResponseFieldLines(t *testing.T) {
	idx := newIdx()
	idx.RequestFields["op"] = map[string]int{"email": 42}
	idx.ResponseFields["op"] = map[string]int{"token": 99}

	req := map[string]FieldConstraint{"email": {Type: "string"}}
	fillRequestFieldLines(req, "op", idx)
	if req["email"].Line != 42 {
		t.Errorf("request line = %d, want 42", req["email"].Line)
	}

	resp := map[string]FieldConstraint{"token": {Type: "string"}}
	fillResponseFieldLines(resp, "op", idx)
	if resp["token"].Line != 99 {
		t.Errorf("response line = %d, want 99", resp["token"].Line)
	}
}

func TestExtractRequestResponseConstraintsOps(t *testing.T) {
	item := &openapi3.PathItem{Post: jsonBodyOp("CreateUser", strProps("email"))}
	result := map[string]map[string]FieldConstraint{}
	extractRequestConstraintsOps(result, item, nil)
	if result["CreateUser"]["email"].Type != "string" {
		t.Errorf("request ops result = %v", result)
	}

	item2 := &openapi3.PathItem{Get: jsonRespOp("GetUser", strProps("id"))}
	result2 := map[string]map[string]FieldConstraint{}
	extractResponseConstraintsOps(result2, item2, nil)
	if result2["GetUser"]["id"].Type != "string" {
		t.Errorf("response ops result = %v", result2)
	}
}
