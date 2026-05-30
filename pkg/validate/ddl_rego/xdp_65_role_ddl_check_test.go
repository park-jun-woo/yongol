//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what XDP-65 test — Rego role 값이 DDL role CHECK 제약에 선언되어야 함

package ddl_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func roleTable(roles ...string) ddl.Table {
	return ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"role": {Name: "role", CheckEnum: roles},
		},
	}
}

func TestXdp65RoleDDLCheck(t *testing.T) {
	if d := xdp65RoleDDLCheck(nil); d != nil {
		t.Errorf("nil fs should yield nil, got %v", d)
	}

	// No role CHECK at all -> rule passes (no role model).
	noRole := &yongol.Fullstack{
		DDLTables: []ddl.Table{ddlTable("users", "id")},
		ParsedPolicies: []rego.Policy{
			{File: "p.rego", Rules: []rego.AllowRule{{UsesRole: true, RoleValue: "admin", SourceLine: 2}}},
		},
	}
	if d := xdp65RoleDDLCheck(noRole); d != nil {
		t.Errorf("no role CHECK should pass, got %v", d)
	}

	// All roles declared -> no diag. UsesRole=false and empty RoleValue skipped.
	ok := &yongol.Fullstack{
		DDLTables: []ddl.Table{roleTable("admin", "user")},
		ParsedPolicies: []rego.Policy{
			{File: "p.rego", Rules: []rego.AllowRule{
				{UsesRole: true, RoleValue: "admin", SourceLine: 2},
				{UsesRole: true, RoleValue: "admin", SourceLine: 9}, // dup -> first wins
				{UsesRole: false, RoleValue: "ignored"},             // not a role use
				{UsesRole: true, RoleValue: ""},                     // empty
			}},
		},
	}
	if d := xdp65RoleDDLCheck(ok); len(d) != 0 {
		t.Errorf("expected no diags, got %v", d)
	}

	// Undeclared role -> diag.
	bad := &yongol.Fullstack{
		DDLTables: []ddl.Table{roleTable("admin")},
		ParsedPolicies: []rego.Policy{
			{File: "p.rego", Rules: []rego.AllowRule{{UsesRole: true, RoleValue: "superuser", SourceLine: 4}}},
		},
	}
	d := xdp65RoleDDLCheck(bad)
	if len(d) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDP-65]") || !strings.Contains(d[0].Message, "superuser") || d[0].Line != 4 {
		t.Errorf("unexpected diag: %+v", d[0])
	}
}
