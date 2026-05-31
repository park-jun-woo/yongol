//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestIndexSchemaEntryDirect(t *testing.T) {
	doc := docNode(t, `User:
  properties:
    id:
      type: integer
`)
	nameKey := doc.Content[0]
	schemaNode := doc.Content[1]
	idx := newIdx()
	indexSchemaEntry(nameKey, schemaNode, idx)
	if idx.Schemas["User"] == 0 {
		t.Errorf("schema not indexed: %v", idx.Schemas)
	}
	if idx.SchemaProperties["User"]["id"] == 0 {
		t.Errorf("schema property not indexed: %v", idx.SchemaProperties)
	}
	// schema with no properties -> only schema line recorded
	doc2 := docNode(t, `Empty:
  type: object
`)
	idx2 := newIdx()
	indexSchemaEntry(doc2.Content[0], doc2.Content[1], idx2)
	if idx2.Schemas["Empty"] == 0 {
		t.Errorf("schema line should still be recorded")
	}
	if _, ok := idx2.SchemaProperties["Empty"]; ok {
		t.Errorf("no properties should mean no SchemaProperties entry")
	}
}
