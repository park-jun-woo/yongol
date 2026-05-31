//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestByName_ZeroCov — openapi 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestByNameWalkSchemas_ZeroCov(t *testing.T) {
	const doc = `Item:
  type: object
  properties:
    id:
      type: integer
    name:
      type: string
`
	var root yaml.Node
	if err := yaml.Unmarshal([]byte(doc), &root); err != nil {
		t.Fatal(err)
	}
	schemas := root.Content[0]
	idx := newLineIndexForTest()
	walkSchemas(schemas, idx)
	if len(idx.Schemas) == 0 {
		t.Errorf("walkSchemas indexed no schemas")
	}

	// non-mapping node short-circuits.
	var scalar yaml.Node
	_ = yaml.Unmarshal([]byte("scalar\n"), &scalar)
	walkSchemas(scalar.Content[0], idx)
}
