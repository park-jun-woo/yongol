//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what xqp30OwnerLookupQuery — ownership 쿼리 존재 검증 (nil/빈/pass/fire) 검증

package query_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqp30OwnerLookupQuery(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := xqp30OwnerLookupQuery(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no policies returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xqp30OwnerLookupQuery(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("matching query passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SQLcQueries: []sqlc.QuerySpec{
				{Name: "OwnerLookupOrder"},
			},
			ParsedPolicies: []rego.Policy{
				{
					File: "auth.rego",
					Ownerships: []rego.OwnershipMapping{
						{Resource: "order", Table: "orders", Column: "user_id"},
					},
				},
			},
		}
		diags := xqp30OwnerLookupQuery(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing query fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ParsedPolicies: []rego.Policy{
				{
					File: "auth.rego",
					Ownerships: []rego.OwnershipMapping{
						{Resource: "order", Table: "orders", Column: "user_id", SourceLine: 5},
					},
				},
			},
		}
		diags := xqp30OwnerLookupQuery(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XQP-30]") {
			t.Errorf("expected XQP-30, got %s", diags[0].Message)
		}
	})
}
