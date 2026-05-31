//ff:func feature=manifest type=test control=sequence
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증
package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestWalkPathsNonMapping(t *testing.T) {
	scalar := &yaml.Node{Kind: yaml.ScalarNode, Value: "x"}
	idx := newIdx()
	walkPaths(scalar, idx) // must not panic, no-op
	if len(idx.Paths) != 0 {
		t.Errorf("expected empty paths")
	}
}
