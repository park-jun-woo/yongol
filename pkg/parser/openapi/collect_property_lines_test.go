//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestCollectPropertyLines(t *testing.T) {
	doc := docNode(t, `properties:
  a:
    type: string
  b:
    type: integer
`)
	props := mapValue(doc, "properties")
	got := collectPropertyLines(props)
	if got["a"] == 0 || got["b"] == 0 {
		t.Errorf("collectPropertyLines = %v", got)
	}
	// nil / non-mapping returns empty map
	if g := collectPropertyLines(nil); len(g) != 0 {
		t.Errorf("nil props = %v", g)
	}
}
