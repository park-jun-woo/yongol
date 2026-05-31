//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestExtractReturningClause(t *testing.T) {
	tests := []struct{ body, want string }{
		{"INSERT INTO users VALUES (1) RETURNING *;", "*"},
		{"INSERT INTO users VALUES (1) RETURNING id, email", "id, email"},
		{"SELECT * FROM users WHERE id = @id;", ""},
		{"", ""},
		{"INSERT ... RETURNING id; -- partial", "id"},
	}
	for _, tt := range tests {
		if got := extractReturningClause(tt.body); got != tt.want {
			t.Errorf("extractReturningClause(%q) = %q, want %q", tt.body, got, tt.want)
		}
	}
}
