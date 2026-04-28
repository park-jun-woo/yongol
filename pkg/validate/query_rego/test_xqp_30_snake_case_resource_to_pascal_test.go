//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestXqp30_Snake_Case_Resource_To_Pascal_Query_Name — snake_case→Pascal 매핑 검증

package query_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqp30_Snake_Case_Resource_To_Pascal_Query_Name(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{
				File: "policy.rego",
				Ownerships: []rego.OwnershipMapping{
					{Resource: "execution_log", Table: "execution_logs", Column: "org_id"},
				},
			},
		},
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "OwnerLookupExecutionLog"},
		},
	}
	if diags := xqp30OwnerLookupQuery(fs); len(diags) != 0 {
		t.Errorf("expected no diagnostics for matching snake→Pascal, got: %+v", diags)
	}
}
