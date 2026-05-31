//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestWalkComponentsAndSchemas(t *testing.T) {
	doc := docNode(t, `components:
  schemas:
    User:
      properties:
        id:
          type: integer
        name:
          type: string
`)
	idx := newIdx()
	comps := mapValue(doc, "components")
	walkComponents(comps, idx)

	if idx.Schemas["User"] == 0 {
		t.Errorf("schema User not indexed: %v", idx.Schemas)
	}
	props := idx.SchemaProperties["User"]
	if props["id"] == 0 || props["name"] == 0 {
		t.Errorf("schema properties not indexed: %v", props)
	}
}
