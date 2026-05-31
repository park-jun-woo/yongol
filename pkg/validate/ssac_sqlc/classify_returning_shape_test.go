//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestClassifyReturningShape(t *testing.T) {
	tbl := ddlTable("id", "email")
	tests := []struct {
		name   string
		clause string
		table  *ddl.Table
		want   ReturningShape
	}{
		{"empty", "", tbl, ShapeNone},
		{"star", "*", tbl, ShapeFull},
		{"covers all", "id, email", tbl, ShapeFull},
		{"partial subset", "id", tbl, ShapePartial},
		{"nil table non-empty", "id", nil, ShapePartial},
		{"empty-column table", "id", &ddl.Table{}, ShapePartial},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classifyReturningShape(tt.clause, tt.table); got != tt.want {
				t.Errorf("classifyReturningShape(%q) = %q, want %q", tt.clause, got, tt.want)
			}
		})
	}
}
