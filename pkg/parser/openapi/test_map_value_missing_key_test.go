//ff:func feature=openapi-parse type=test control=sequence
//ff:what mapValue — 존재하지 않는 key 는 nil, 존재 key 는 값 반환

package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMapValue_MissingKey(t *testing.T) {
	// Build a mapping with key=foo
	m := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "foo"},
		{Kind: yaml.ScalarNode, Value: "bar"},
	}}
	if got := mapValue(m, "missing"); got != nil {
		t.Errorf("missing key should return nil, got %v", got)
	}
	got := mapValue(m, "foo")
	if got == nil || got.Value != "bar" {
		t.Errorf("present key got %v, want bar", got)
	}
}
