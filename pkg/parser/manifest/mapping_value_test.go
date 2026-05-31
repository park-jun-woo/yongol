//ff:func feature=manifest type=test control=sequence
//ff:what TestManifestHelpers — unit tests for the pure manifest parser helper functions
package manifest

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMappingValue(t *testing.T) {
	var root yaml.Node
	if err := yaml.Unmarshal([]byte("a: 1\nb: hello\n"), &root); err != nil {
		t.Fatal(err)
	}
	doc := root.Content[0] // document → mapping node
	if v := mappingValue(doc, "b"); v == nil || v.Value != "hello" {
		t.Errorf("mappingValue(b) = %v", v)
	}
	if v := mappingValue(doc, "missing"); v != nil {
		t.Errorf("missing key should return nil, got %v", v)
	}
	// nil node → nil.
	if v := mappingValue(nil, "a"); v != nil {
		t.Error("nil node → nil")
	}
	// non-mapping node → nil.
	scalar := &yaml.Node{Kind: yaml.ScalarNode, Value: "x"}
	if v := mappingValue(scalar, "a"); v != nil {
		t.Error("scalar node → nil")
	}
}
