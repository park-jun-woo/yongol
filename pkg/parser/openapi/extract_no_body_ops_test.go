//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractNoBodyOps — requestBody 유무에 따라 올바른 집합을 반환하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractNoBodyOps(t *testing.T) {
	bodySchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	doc := &openapi3.T{
		Paths: &openapi3.Paths{},
	}
	doc.Paths.Set("/rooms", &openapi3.PathItem{
		Post: &openapi3.Operation{
			OperationID: "CreateRoom",
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{Value: bodySchema},
						},
					},
				},
			},
		},
		Get: &openapi3.Operation{
			OperationID: "ListRooms",
		},
	})
	doc.Paths.Set("/rooms/{roomId}", &openapi3.PathItem{
		Delete: &openapi3.Operation{
			OperationID: "DeleteRoom",
		},
	})

	result := ExtractNoBodyOps(doc)

	if result["CreateRoom"] {
		t.Error("CreateRoom has requestBody, should not be in NoBodyOps")
	}
	if !result["ListRooms"] {
		t.Error("ListRooms has no requestBody, should be in NoBodyOps")
	}
	if !result["DeleteRoom"] {
		t.Error("DeleteRoom has no requestBody, should be in NoBodyOps")
	}
}
