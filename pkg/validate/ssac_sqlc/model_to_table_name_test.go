//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestModelToTableName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"User", "users"},
		{"AuditLog", "audit_logs"},
		{"Workflow", "workflows"},
	}
	for _, tt := range tests {
		if got := modelToTableName(tt.in); got != tt.want {
			t.Errorf("modelToTableName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
