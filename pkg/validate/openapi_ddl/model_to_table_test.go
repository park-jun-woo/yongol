//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what modelToTable — PascalCase -> snake_case + plural 변환 검증

package openapi_ddl

import "testing"

func TestModelToTable(t *testing.T) {
	tests := []struct {
		model string
		want  string
	}{
		{"User", "users"},
		{"Order", "orders"},
		{"WorkflowAction", "workflow_actions"},
		{"Category", "categories"},
		{"users", "users"},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			got := modelToTable(tt.model)
			if got != tt.want {
				t.Errorf("modelToTable(%q) = %q, want %q", tt.model, got, tt.want)
			}
		})
	}
}
