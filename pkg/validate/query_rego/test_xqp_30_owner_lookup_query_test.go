//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what XQP-30 — @ownership 매핑에 OwnerLookup<Resource> sqlc 쿼리 존재 강제

package query_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
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

func TestXqp30_Via_Advice_Uses_Join_Shape(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{
				File: "policy.rego",
				Ownerships: []rego.OwnershipMapping{
					{
						Resource:  "lesson",
						Table:     "courses",
						Column:    "instructor_id",
						JoinTable: "lessons",
						JoinFK:    "course_id",
					},
				},
			},
		},
	}
	diags := xqp30OwnerLookupQuery(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Advice, "JOIN lessons l ON l.course_id = c.id") {
		t.Errorf("advice missing JOIN shape: %s", diags[0].Advice)
	}
}
