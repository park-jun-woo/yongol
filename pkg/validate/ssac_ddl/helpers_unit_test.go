//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	parserddl "github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsPlural(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"Workflows", true},
		{"Workflow", false},
		{"users", true},
		{"user", false},
		{"", false},
	}
	for _, tt := range tests {
		if got := isPlural(tt.in); got != tt.want {
			t.Errorf("isPlural(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestNormalizeTypeName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"[]Reservation", "Reservation"},
		{"*User", "User"},
		{"User", "User"},
		{"[]*Order", "Order"}, // strips slice prefix, then pointer prefix
	}
	for _, tt := range tests {
		if got := normalizeTypeName(tt.in); got != tt.want {
			t.Errorf("normalizeTypeName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

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

func TestIsAuthRequiredTable(t *testing.T) {
	// Backend.Auth configured → refresh_tokens is auth-required.
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	fs.Manifest.Backend.Auth = &manifest.Auth{}
	if !isAuthRequiredTable(fs, "refresh_tokens") {
		t.Error("refresh_tokens should be auth-required when auth configured")
	}
	if isAuthRequiredTable(fs, "courses") {
		t.Error("courses should not be auth-required")
	}
	// no manifest → false.
	if isAuthRequiredTable(&yongol.Fullstack{}, "refresh_tokens") {
		t.Error("nil manifest → not auth-required")
	}
	// manifest without Backend.Auth → false.
	if isAuthRequiredTable(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}, "refresh_tokens") {
		t.Error("no Backend.Auth → not auth-required")
	}
}

func TestDDLTableSet(t *testing.T) {
	// Ground lookup path.
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.table": {"users": true, "courses": true},
	}})
	set := ddlTableSet(fs)
	if !set["users"] || !set["courses"] {
		t.Errorf("ground lookup set = %v", set)
	}

	// Fallback path (no Ground): build from DDLTables.
	fs2 := &yongol.Fullstack{DDLTables: []parserddl.Table{{Name: "orders"}}}
	set2 := ddlTableSet(fs2)
	if !set2["orders"] {
		t.Errorf("fallback set = %v", set2)
	}
}

func TestDDLColumnSet(t *testing.T) {
	// Ground lookup path.
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.column.users": {"id": true, "email": true},
	}})
	set, ok := ddlColumnSet(fs, "users")
	if !ok || !set["email"] {
		t.Errorf("ground column set = %v ok=%v", set, ok)
	}

	// Fallback path.
	fs2 := &yongol.Fullstack{DDLTables: []parserddl.Table{
		{Name: "orders", Columns: map[string]parserddl.Column{"id": {Name: "id"}}},
	}}
	set2, ok := ddlColumnSet(fs2, "orders")
	if !ok || !set2["id"] {
		t.Errorf("fallback column set = %v ok=%v", set2, ok)
	}
	// Unknown table → (nil,false).
	if _, ok := ddlColumnSet(fs2, "missing"); ok {
		t.Error("missing table should return ok=false")
	}
}

func TestIsArchivedTable(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Flags: rule.StringSet{"archived.old_logs": true}})
	if !isArchivedTable(fs, "old_logs") {
		t.Error("old_logs should be archived")
	}
	if isArchivedTable(fs, "users") {
		t.Error("users not archived")
	}
	// nil Ground → false.
	if isArchivedTable(&yongol.Fullstack{}, "old_logs") {
		t.Error("nil Ground → false")
	}
}

func TestIsPkgModelTable(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Flags: rule.StringSet{"pkgModel.sessions": true}})
	if !isPkgModelTable(fs, "sessions") {
		t.Error("sessions should be a pkg model table")
	}
	if isPkgModelTable(fs, "users") {
		t.Error("users not a pkg model table")
	}
	if isPkgModelTable(&yongol.Fullstack{}, "sessions") {
		t.Error("nil Ground → false")
	}
}

func TestCanonicalDDLTableSet(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.table": {"Workflows": true, "AuditLogs": true},
	}})
	set := canonicalDDLTableSet(fs)
	// canonicalTableKey lowercases + snake + singular.
	if !set["workflow"] {
		t.Errorf("expected canonical 'workflow' key, got %v", set)
	}
	if !set["audit_log"] {
		t.Errorf("expected canonical 'audit_log' key, got %v", set)
	}
}
