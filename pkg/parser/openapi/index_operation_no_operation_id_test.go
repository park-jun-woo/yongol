//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestIndexOperationNoOperationID(t *testing.T) {
	doc := docNode(t, `summary: no id here
`)
	idx := newIdx()
	indexOperation(doc, idx)
	if len(idx.Operations) != 0 {
		t.Errorf("expected no operations, got %v", idx.Operations)
	}
}
