//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what checkOwnershipMapping — ownership 쿼리 존재 검증 (pass/fire/empty/duplicate) 검증

package query_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestCheckOwnershipMapping(t *testing.T) {
	t.Run("empty resource skips", func(t *testing.T) {
		om := rego.OwnershipMapping{Resource: ""}
		_, fired := checkOwnershipMapping("auth.rego", om, nil, make(map[string]bool))
		if fired {
			t.Error("expected false for empty resource")
		}
	})

	t.Run("query present passes", func(t *testing.T) {
		om := rego.OwnershipMapping{Resource: "order", Table: "orders", Column: "user_id"}
		have := map[string]bool{"OwnerLookupOrder": true}
		_, fired := checkOwnershipMapping("auth.rego", om, have, make(map[string]bool))
		if fired {
			t.Error("expected false when query exists")
		}
	})

	t.Run("query missing fires XQP-30", func(t *testing.T) {
		om := rego.OwnershipMapping{Resource: "order", Table: "orders", Column: "user_id", SourceLine: 10}
		have := map[string]bool{}
		diag, fired := checkOwnershipMapping("auth.rego", om, have, make(map[string]bool))
		if !fired {
			t.Fatal("expected true")
		}
		if !strings.Contains(diag.Message, "[XQP-30]") {
			t.Errorf("expected XQP-30, got %s", diag.Message)
		}
		if !strings.Contains(diag.Message, "OwnerLookupOrder") {
			t.Errorf("expected query name, got %s", diag.Message)
		}
	})

	t.Run("duplicate same file skips", func(t *testing.T) {
		om := rego.OwnershipMapping{Resource: "order", Table: "orders", Column: "user_id"}
		have := map[string]bool{}
		seen := make(map[string]bool)
		// First call fires
		_, fired1 := checkOwnershipMapping("auth.rego", om, have, seen)
		if !fired1 {
			t.Fatal("first call expected to fire")
		}
		// Second call skips (duplicate)
		_, fired2 := checkOwnershipMapping("auth.rego", om, have, seen)
		if fired2 {
			t.Error("expected false for duplicate")
		}
	})

	t.Run("snake_case resource converts to PascalCase", func(t *testing.T) {
		om := rego.OwnershipMapping{Resource: "execution_log", Table: "execution_logs", Column: "user_id"}
		have := map[string]bool{}
		diag, fired := checkOwnershipMapping("auth.rego", om, have, make(map[string]bool))
		if !fired {
			t.Fatal("expected true")
		}
		if !strings.Contains(diag.Message, "OwnerLookupExecutionLog") {
			t.Errorf("expected PascalCase name, got %s", diag.Message)
		}
	})
}
