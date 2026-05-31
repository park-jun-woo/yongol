//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestIndexRequestBodyDirect(t *testing.T) {
	doc := docNode(t, `content:
  application/json:
    schema:
      properties:
        name:
          type: string
`)
	idx := newIdx()
	indexRequestBody(doc, "op1", idx)
	if idx.RequestFields["op1"]["name"] == 0 {
		t.Errorf("request body field not indexed: %v", idx.RequestFields)
	}
	// body without schema props -> no-op
	empty := docNode(t, `content: {}
`)
	idx2 := newIdx()
	indexRequestBody(empty, "op2", idx2)
	if len(idx2.RequestFields) != 0 {
		t.Errorf("expected no request fields")
	}
}
