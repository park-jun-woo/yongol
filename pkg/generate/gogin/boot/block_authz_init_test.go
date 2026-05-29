//ff:func feature=gen-gogin type=test control=iteration dimension=2
//ff:what blockAuthzInit — OPA authz.Init(policyPath, ownerships) — DB 의존 없음

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockAuthzInit_OwnershipMappings(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{Ownerships: []rego.OwnershipMapping{
				{Resource: "note", Table: "notes", Column: "owner_id"},
				{Resource: "comment", Table: "comments", Column: "user_id", JoinTable: "notes", JoinFK: "note_id"},
			}},
		},
	}
	block := blockAuthzInit(fs)
	if block.Name != "authz-init" {
		t.Errorf("name = %q, want authz-init", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `initAuthz(os.Getenv("OPA_POLICY_PATH"))`) {
		t.Errorf("must call initAuthz with OPA_POLICY_PATH, got:\n%s", body)
	}
	factory := strings.Join(block.Funcs, "\n")
	if !strings.Contains(factory, `{Resource: "note", Table: "notes", Column: "owner_id"},`) {
		t.Errorf("ownership mapping missing in factory, got:\n%s", factory)
	}
	if !strings.Contains(factory, `JoinTable: "notes", JoinFK: "note_id"`) {
		t.Errorf("join ownership mapping missing in factory, got:\n%s", factory)
	}
}

func TestBlockAuthzInit_ActiveOnAuthSequence(t *testing.T) {
	// Active predicate must be hasAuthSequence (gated, not always active).
	block := blockAuthzInit(&yongol.Fullstack{})
	if block.Active == nil {
		t.Fatalf("authz-init must carry an Active predicate")
	}
	if block.Active(&yongol.Fullstack{}) {
		t.Errorf("authz-init must be inactive without an @auth sequence")
	}
}
