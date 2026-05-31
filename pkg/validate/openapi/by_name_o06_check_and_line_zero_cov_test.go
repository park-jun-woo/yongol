//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what TestByName_ZeroCov — O-6 스키마 워커들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"testing"
)

func TestByNameO06CheckAndLine_ZeroCov(t *testing.T) {
	doc := byNameO06Doc()
	fs := byNameO06FS(doc)

	entry := o06SchemaEntry{
		schema:     doc.Components.Schemas["Workflow"].Value,
		schemaName: "Workflow",
	}
	diags := o06CheckSchemaRequired(fs, entry)
	// "missing" is required but not a property → a diagnostic is expected.
	if len(diags) == 0 {
		t.Errorf("o06CheckSchemaRequired expected diagnostics for dangling required")
	}

	if line := o06RequiredLine(fs, "Workflow", "name"); line == 0 {
		t.Errorf("o06RequiredLine fallback returned 0")
	}
	_ = o06RequiredLine(fs, "Unknown", "x")
}
