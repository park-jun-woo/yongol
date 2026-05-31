//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestSchemaPropsOfBody(t *testing.T) {
	body := parseNode(t, respYAML)
	props := schemaPropsOfBody(body)
	if props == nil || props.Kind != yaml.MappingNode {
		t.Fatalf("expected properties mapping, got %v", props)
	}
	lines := collectPropertyLines(props)
	if _, ok := lines["id"]; !ok {
		t.Errorf("missing id in property lines: %v", lines)
	}

	// Missing content → nil.
	noContent := parseNode(t, "summary: hello\n")
	if got := schemaPropsOfBody(noContent); got != nil {
		t.Errorf("no content → %v, want nil", got)
	}
}
