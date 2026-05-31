//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestWalkPathsAndOperations(t *testing.T) {
	doc := docNode(t, `paths:
  /login:
    post:
      operationId: login
      requestBody:
        content:
          application/json:
            schema:
              properties:
                email:
                  type: string
      responses:
        "200":
          content:
            application/json:
              schema:
                properties:
                  token:
                    type: string
`)
	idx := newIdx()
	paths := mapValue(doc, "paths")
	walkPaths(paths, idx)

	if idx.Paths["/login"] == 0 {
		t.Errorf("path /login not indexed: %v", idx.Paths)
	}
	if idx.Operations["login"] == 0 {
		t.Errorf("operation login not indexed: %v", idx.Operations)
	}
	if idx.RequestFields["login"]["email"] == 0 {
		t.Errorf("request field email not indexed: %v", idx.RequestFields)
	}
	if idx.ResponseFields["login"]["token"] == 0 {
		t.Errorf("response field token not indexed: %v", idx.ResponseFields)
	}
}
