//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestIndexPathItemDirect(t *testing.T) {
	doc := docNode(t, `/users/{id}:
  get:
    operationId: getUser
    responses:
      "200":
        content:
          application/json:
            schema:
              properties:
                id:
                  type: integer
`)
	// doc is the mapping; its first key/value pair is the path entry.
	pathKey := doc.Content[0]
	pathItem := doc.Content[1]
	idx := newIdx()
	indexPathItem(pathKey, pathItem, idx)
	if idx.Paths["/users/{id}"] == 0 || idx.Operations["getUser"] == 0 {
		t.Errorf("indexPathItem failed: paths=%v ops=%v", idx.Paths, idx.Operations)
	}
	if idx.ResponseFields["getUser"]["id"] == 0 {
		t.Errorf("response field not indexed: %v", idx.ResponseFields)
	}
}
