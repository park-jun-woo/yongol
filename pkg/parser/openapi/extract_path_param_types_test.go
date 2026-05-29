//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractPathParamTypes — path 파라미터의 타입을 올바르게 추출하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractPathParamTypes(t *testing.T) {
	doc := &openapi3.T{
		Paths: &openapi3.Paths{},
	}
	doc.Paths.Set("/buildings/{id}", &openapi3.PathItem{
		Get: &openapi3.Operation{
			OperationID: "GetBuilding",
			Parameters: openapi3.Parameters{
				&openapi3.ParameterRef{
					Value: &openapi3.Parameter{
						Name: "id",
						In:   "path",
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: &openapi3.Types{"integer"},
							},
						},
					},
				},
			},
		},
		Delete: &openapi3.Operation{
			OperationID: "DeleteBuilding",
			Parameters: openapi3.Parameters{
				&openapi3.ParameterRef{
					Value: &openapi3.Parameter{
						Name: "id",
						In:   "path",
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: &openapi3.Types{"integer"},
							},
						},
					},
				},
			},
		},
	})
	doc.Paths.Set("/buildings/{id}/units", &openapi3.PathItem{
		Get: &openapi3.Operation{
			OperationID: "ListUnits",
			Parameters: openapi3.Parameters{
				&openapi3.ParameterRef{
					Value: &openapi3.Parameter{
						Name: "id",
						In:   "path",
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: &openapi3.Types{"string"},
							},
						},
					},
				},
			},
		},
	})

	result := ExtractPathParamTypes(doc)

	if result["GetBuilding"]["id"] != "integer" {
		t.Errorf("GetBuilding.id = %q, want \"integer\"", result["GetBuilding"]["id"])
	}
	if result["DeleteBuilding"]["id"] != "integer" {
		t.Errorf("DeleteBuilding.id = %q, want \"integer\"", result["DeleteBuilding"]["id"])
	}
	if result["ListUnits"]["id"] != "string" {
		t.Errorf("ListUnits.id = %q, want \"string\"", result["ListUnits"]["id"])
	}
}
