//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what isColumnNotNull — NOT NULL 제약/PK/nullable 컬럼 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsColumnNotNull(t *testing.T) {
	tbl := &ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":    {NotNull: false}, // PK but not explicitly NOT NULL
			"email": {NotNull: true},
			"phone": {NotNull: false},
			"bio":   {},
		},
		PrimaryKey: []string{"id"},
	}

	tests := []struct {
		name string
		col  string
		want bool
	}{
		{"NOT NULL constraint", "email", true},
		{"primary key column (implicit NOT NULL)", "id", true},
		{"nullable column", "phone", false},
		{"default nullable", "bio", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isColumnNotNull(tbl, tt.col)
			if got != tt.want {
				t.Errorf("isColumnNotNull(%q) = %v, want %v", tt.col, got, tt.want)
			}
		})
	}
}
