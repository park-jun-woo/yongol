//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestIndexOperationDirect(t *testing.T) {
	doc := docNode(t, `operationId: createUser
requestBody:
  content:
    application/json:
      schema:
        properties:
          email:
            type: string
responses:
  "201":
    content:
      application/json:
        schema:
          properties:
            id:
              type: integer
`)
	idx := newIdx()
	indexOperation(doc, idx)
	if idx.Operations["createUser"] == 0 {
		t.Errorf("operation not indexed: %v", idx.Operations)
	}
	if idx.RequestFields["createUser"]["email"] == 0 {
		t.Errorf("request field not indexed: %v", idx.RequestFields)
	}
	if idx.ResponseFields["createUser"]["id"] == 0 {
		t.Errorf("response field not indexed: %v", idx.ResponseFields)
	}
}
