//ff:func feature=gen-gogin type=test control=sequence topic=authz
//ff:what TestBuildAuth_ResourceIDAbsent_SkipsOwnerLookup — ResourceID 부재 시 skip (BUG-033)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildAuth_ResourceIDAbsent_SkipsOwnerLookup verifies that when
// the `@auth` sequence has no ResourceID input at all (creation-form
// endpoints like POST /workflows), buildAuth skips the OwnerLookup
// injection and passes Owners: nil. Previously the emitter fell back
// to `ctx, 0` which always yielded sql.ErrNoRows → 403. (Phase005 /
// BUG-033)
func TestBuildAuth_ResourceIDAbsent_SkipsOwnerLookup(t *testing.T) {
	g := &methodGen{
		FuncName:   "CreateWorkflow",
		FileName:   "workflow_service.ssac",
		ModulePath: "example.com/zenflow",
		UseTx:      true,
		FirstErr:   true,
		Ownerships: []rego.OwnershipMapping{
			{Resource: "workflow", Table: "workflows", Column: "owner_id"},
		},
		DeclaredVars: make(map[string]bool),
	}
	seq := ssacparser.Sequence{
		Type:     "auth",
		Action:   "CreateWorkflow",
		Resource: "workflow",
		// Inputs empty — no ResourceID key at all.
	}
	lines, _ := g.buildAuth(seq)
	body := strings.Join(lines, "\n")

	if strings.Contains(body, "OwnerLookup") {
		t.Fatalf("ResourceID absent: must not emit OwnerLookup line, got:\n%s", body)
	}
	if !strings.Contains(body, "Owners: nil") {
		t.Fatalf("ResourceID absent: Owners must be nil, got:\n%s", body)
	}
}
