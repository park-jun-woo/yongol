//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsSqlcRowType(t *testing.T) {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "UserFindByEmail", RowType: "UserFindByEmailRow"},
			{Name: "ListUsers", RowType: ""},
		},
	}
	if !isSqlcRowType(fs, "UserFindByEmailRow") {
		t.Error("expected UserFindByEmailRow to be a sqlc row type")
	}
	if isSqlcRowType(fs, "Unknown") {
		t.Error("Unknown should not match")
	}
	// nil fs / empty typeName guards.
	if isSqlcRowType(nil, "X") {
		t.Error("nil fs should yield false")
	}
	if isSqlcRowType(fs, "") {
		t.Error("empty typeName should yield false")
	}
}
