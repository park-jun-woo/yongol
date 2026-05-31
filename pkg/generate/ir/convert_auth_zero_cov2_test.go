//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestConvertAuth_ZeroCov2(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{Ownerships: []rego.OwnershipMapping{{Resource: "project", Table: "projects", Column: "owner_id"}}},
		},
		DDLTables: []ddl.Table{{Name: "projects", PrimaryKey: []string{"id"}}},
	}
	// ownership populated
	op := convertAuth(ssac.Sequence{Action: "delete", Resource: "project", Inputs: map[string]string{"ResourceID": "project.ID"}}, fs)
	if op.Kind != OpAuth || op.Auth == nil {
		t.Fatalf("expected OpAuth, got %+v", op)
	}
	if op.Auth.StatusCode != 403 {
		t.Errorf("default status = %d", op.Auth.StatusCode)
	}
	if op.Auth.Ownership == nil || op.Auth.Ownership.Table != "projects" || op.Auth.Ownership.ResourcePK != "id" {
		t.Errorf("ownership = %+v", op.Auth.Ownership)
	}
	// zero ResourceID → no ownership; custom status
	op = convertAuth(ssac.Sequence{Action: "read", Resource: "project", ErrStatus: 401, Inputs: map[string]string{"ResourceID": "0"}}, fs)
	if op.Auth.Ownership != nil || op.Auth.StatusCode != 401 {
		t.Errorf("expected no ownership, status 401, got %+v / %d", op.Auth.Ownership, op.Auth.StatusCode)
	}
	// no matching policy resource
	op = convertAuth(ssac.Sequence{Resource: "other", Inputs: map[string]string{"ResourceID": "x.ID"}}, fs)
	if op.Auth.Ownership != nil {
		t.Errorf("expected no ownership for unmatched resource")
	}
	// nil fs
	op = convertAuth(ssac.Sequence{Resource: "project", Inputs: map[string]string{"ResourceID": "x.ID"}}, nil)
	if op.Auth.Ownership != nil {
		t.Errorf("expected no ownership for nil fs")
	}
}
