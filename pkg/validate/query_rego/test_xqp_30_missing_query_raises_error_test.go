//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestXqp30_Missing_Query_Raises_Error — @ownership 대응 쿼리 부재 시 [XQP-30]

package query_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqp30_Missing_Query_Raises_Error(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{
				File: "policy.rego",
				Ownerships: []rego.OwnershipMapping{
					{Resource: "workflow", Table: "workflows", Column: "org_id", SourceLine: 3},
				},
			},
		},
		SQLcQueries: nil, // no queries provided
	}
	diags := xqp30OwnerLookupQuery(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "OwnerLookupWorkflow") {
		t.Errorf("message missing expected query name: %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Advice, "OwnerLookupWorkflow") {
		t.Errorf("advice missing query stub: %s", diags[0].Advice)
	}
}
