//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestXqp30_Present_Query_No_Diagnostic — 쿼리 존재 시 진단 없음

package query_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqp30_Present_Query_No_Diagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{
				File: "policy.rego",
				Ownerships: []rego.OwnershipMapping{
					{Resource: "workflow", Table: "workflows", Column: "org_id"},
				},
			},
		},
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "OwnerLookupWorkflow"},
		},
	}
	if diags := xqp30OwnerLookupQuery(fs); len(diags) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", diags)
	}
}
