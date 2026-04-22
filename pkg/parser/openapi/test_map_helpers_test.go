//ff:func feature=manifest type=test control=sequence
//ff:what mapKey / mapValue nil + non-mapping 경로 회귀

package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMapKey_NilAndNonMapping(t *testing.T) {
	if got := mapKey(nil, "x"); got != nil {
		t.Errorf("nil in, got %v", got)
	}
	scalar := &yaml.Node{Kind: yaml.ScalarNode}
	if got := mapKey(scalar, "x"); got != nil {
		t.Errorf("scalar in, got %v", got)
	}
}

func TestMapValue_NilAndNonMapping(t *testing.T) {
	if got := mapValue(nil, "x"); got != nil {
		t.Errorf("nil in, got %v", got)
	}
	scalar := &yaml.Node{Kind: yaml.ScalarNode}
	if got := mapValue(scalar, "x"); got != nil {
		t.Errorf("scalar in, got %v", got)
	}
}

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
