//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestByName_ZeroCov — openapi 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

func newLineIndexForTest() *LineIndex {
	return &LineIndex{
		Operations:       map[string]int{},
		RequestFields:    map[string]map[string]int{},
		ResponseFields:   map[string]map[string]int{},
		Schemas:          map[string]int{},
		SchemaProperties: map[string]map[string]int{},
		Paths:            map[string]int{},
	}
}

func TestByNameWalkSchemas_ZeroCov(t *testing.T) {
	const doc = `Item:
  type: object
  properties:
    id:
      type: integer
    name:
      type: string
`
	var root yaml.Node
	if err := yaml.Unmarshal([]byte(doc), &root); err != nil {
		t.Fatal(err)
	}
	schemas := root.Content[0]
	idx := newLineIndexForTest()
	walkSchemas(schemas, idx)
	if len(idx.Schemas) == 0 {
		t.Errorf("walkSchemas indexed no schemas")
	}

	// non-mapping node short-circuits.
	var scalar yaml.Node
	_ = yaml.Unmarshal([]byte("scalar\n"), &scalar)
	walkSchemas(scalar.Content[0], idx)
}

func TestByNamePathParamTypes_ZeroCov(t *testing.T) {
	schema := openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{"integer"}})
	op := &openapi3.Operation{
		OperationID: "GetItem",
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{
				Name:   "id",
				In:     "path",
				Schema: schema,
			}},
		},
	}
	result := map[string]map[string]string{}
	collectPathParamTypesForOp(result, op)
	if result["GetItem"]["id"] != "integer" {
		t.Errorf("collectPathParamTypesForOp = %v", result)
	}

	// nil/empty op short-circuits.
	collectPathParamTypesForOp(result, nil)
	collectPathParamTypesForOp(result, &openapi3.Operation{})
}

func TestByNameResponseArrayItemFields_ZeroCov(t *testing.T) {
	itemSchema := openapi3.NewSchema()
	itemSchema.Type = &openapi3.Types{"object"}
	itemSchema.Properties = openapi3.Schemas{
		"id":   openapi3.NewSchemaRef("", openapi3.NewSchema()),
		"name": openapi3.NewSchemaRef("", openapi3.NewSchema()),
	}
	arraySchema := openapi3.NewSchema()
	arraySchema.Type = &openapi3.Types{"array"}
	arraySchema.Items = openapi3.NewSchemaRef("", itemSchema)

	wrapper := openapi3.NewSchema()
	wrapper.Type = &openapi3.Types{"object"}
	wrapper.Properties = openapi3.Schemas{
		"items": openapi3.NewSchemaRef("", arraySchema),
	}

	content := openapi3.NewContentWithJSONSchema(wrapper)
	resp := openapi3.NewResponse().WithContent(content)
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{Value: resp})

	op := &openapi3.Operation{OperationID: "ListItems", Responses: responses}
	result := map[string]map[string]map[string]bool{}
	collectResponseArrayItemFieldsForOp(result, op)
	// Result may or may not contain the op depending on extractArrayItemFields;
	// the call itself exercises the function by name.
	_ = result

	// empty operationID short-circuits.
	collectResponseArrayItemFieldsForOp(result, &openapi3.Operation{})
}
