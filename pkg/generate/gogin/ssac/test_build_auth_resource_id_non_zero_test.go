//ff:func feature=gen-gogin type=test control=sequence topic=authz
//ff:what TestBuildAuth_ResourceIDNonZero_KeepsOwnerLookup — Update/Get/Delete 시 OwnerLookup 유지

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildAuth_ResourceIDNonZero_KeepsOwnerLookup guards the
// Update/Get/Delete path: a non-zero ResourceID expression (e.g.
// `path.id`, `body.WorkflowID`) must still trigger the OwnerLookup
// injection exactly as before.
func TestBuildAuth_ResourceIDNonZero_KeepsOwnerLookup(t *testing.T) {
	g := &methodGen{
		FuncName:   "GetWorkflow",
		FileName:   "workflow_service.ssac",
		ModulePath: "example.com/zenflow",
		UseTx:      false,
		FirstErr:   true,
		Ownerships: []rego.OwnershipMapping{
			{Resource: "workflow", Table: "workflows", Column: "owner_id"},
		},
	}
	seq := ssacparser.Sequence{
		Type:     "auth",
		Action:   "GetWorkflow",
		Resource: "workflow",
		Inputs: map[string]string{
			"ResourceID": "path.id",
		},
	}
	lines, _ := g.buildAuth(seq)
	body := strings.Join(lines, "\n")

	if !strings.Contains(body, "OwnerLookupWorkflow(ctx,") {
		t.Fatalf("ResourceID non-zero: expected OwnerLookup line, got:\n%s", body)
	}
	if strings.Contains(body, "Owners: nil") {
		t.Fatalf("ResourceID non-zero: Owners must be populated, got:\n%s", body)
	}
}
