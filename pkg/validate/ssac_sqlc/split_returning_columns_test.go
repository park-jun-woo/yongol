//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"reflect"
	"testing"
)

func TestSplitReturningColumns(t *testing.T) {
	tests := []struct {
		name   string
		clause string
		want   map[string]bool
	}{
		{"plain", "id, email", map[string]bool{"id": true, "email": true}},
		{"alias prefix", "u.id, u.email", map[string]bool{"id": true, "email": true}},
		{"AS alias", "id AS user_id, email", map[string]bool{"user_id": true, "email": true}},
		{"uppercase folded", "ID, EMAIL", map[string]bool{"id": true, "email": true}},
		{"quoted ident", "\"Name\"", map[string]bool{"name": true}},
		{"empty entries skipped", "id, ,email", map[string]bool{"id": true, "email": true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitReturningColumns(tt.clause)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitReturningColumns(%q) = %v, want %v", tt.clause, got, tt.want)
			}
		})
	}
}
