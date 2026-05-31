//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestIndexFirst2xxResponseDirect(t *testing.T) {
	doc := docNode(t, `"200":
  content:
    application/json:
      schema:
        properties:
          token:
            type: string
`)
	idx := newIdx()
	indexFirst2xxResponse(doc, "login", idx)
	if idx.ResponseFields["login"]["token"] == 0 {
		t.Errorf("first 2xx response field not indexed: %v", idx.ResponseFields)
	}
}
