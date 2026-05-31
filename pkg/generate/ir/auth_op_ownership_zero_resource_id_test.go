//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestAuthOpOwnership -- AuthOp.Ownership Rego @ownership 이식 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestAuthOpOwnershipZeroResourceID(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{
				Ownerships: []rego.OwnershipMapping{
					{Resource: "workflow", Table: "workflows", Column: "owner_id"},
				},
			},
		},
		DDLTables: []ddl.Table{
			{
				Name: "workflows",
				Columns: map[string]ddl.Column{
					"id":       {Name: "id"},
					"owner_id": {Name: "owner_id"},
				},
				PrimaryKey: []string{"id"},
			},
		},
	}

	for _, zero := range []string{"0", "", "nil", "null"} {
		t.Run(zero, func(t *testing.T) {
			assertOwnershipNilForZeroResourceID(t, fs, zero)
		})
	}
}
