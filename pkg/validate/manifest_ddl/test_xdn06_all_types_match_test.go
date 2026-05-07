//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN06_AllTypesMatch_NoDiag — 5 가지 claim 타입 전부 DDL 정합 → PASS

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN06_AllTypesMatch_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Type:      "jwt",
					UserTable: "users",
					Claims: map[string]pmanifest.ClaimDef{
						"ID":      {Key: "id", GoType: "uuid", Typed: true},
						"Email":   {Key: "email", GoType: "string", Typed: true},
						"OrgID":   {Key: "org_id", GoType: "int64", Typed: true},
						"Age":     {Key: "age", GoType: "int32", Typed: true},
						"IsAdmin": {Key: "is_admin", GoType: "bool", Typed: true},
					},
				},
			},
		},
		DDLTables: []ddl.Table{{
			Name: "users",
			Columns: map[string]ddl.Column{
				"id":       {Name: "id", RawType: "UUID"},
				"email":    {Name: "email", RawType: "TEXT"},
				"org_id":   {Name: "org_id", RawType: "BIGINT"},
				"age":      {Name: "age", RawType: "INTEGER"},
				"is_admin": {Name: "is_admin", RawType: "BOOLEAN"},
			},
		}},
	}
	if d := xdn06ClaimDDLType(fs); len(d) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", d)
	}
}
