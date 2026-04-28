//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestXqp30_Via_Advice_Uses_Join_Shape — via 매핑 advice 가 JOIN 형식

package query_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
