//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestWalkComponentsNoSchemas(t *testing.T) {
	doc := docNode(t, `components:
  responses: {}
`)
	idx := newIdx()
	walkComponents(mapValue(doc, "components"), idx)
	if len(idx.Schemas) != 0 {
		t.Errorf("expected no schemas")
	}
}
