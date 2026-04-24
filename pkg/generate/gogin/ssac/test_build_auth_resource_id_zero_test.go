//ff:func feature=gen-gogin type=test control=sequence topic=authz
//ff:what TestBuildAuth_ResourceIDZero — Phase005 (BUG-033) ResourceID==0 시 OwnerLookup 주입 skip

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

// TestBuildAuth_ResourceIDLiteralZero_SkipsOwnerLookup verifies that
// an explicit `{ResourceID: 0}` literal is treated as the creation-
// form signal, same as absence.
func TestBuildAuth_ResourceIDLiteralZero_SkipsOwnerLookup(t *testing.T) {
	g := &methodGen{
		FuncName:   "CreateWorkflow",
		FileName:   "workflow_service.ssac",
		ModulePath: "example.com/zenflow",
		UseTx:      true,
		FirstErr:   true,
		Ownerships: []rego.OwnershipMapping{
			{Resource: "workflow", Table: "workflows", Column: "owner_id"},
		},
	}
	seq := ssacparser.Sequence{
		Type:     "auth",
		Action:   "CreateWorkflow",
		Resource: "workflow",
		Inputs: map[string]string{
			"ResourceID": "0",
		},
	}
	lines, _ := g.buildAuth(seq)
	body := strings.Join(lines, "\n")

	if strings.Contains(body, "OwnerLookup") {
		t.Fatalf("ResourceID=0: must not emit OwnerLookup line, got:\n%s", body)
	}
	if !strings.Contains(body, "Owners: nil") {
		t.Fatalf("ResourceID=0: Owners must be nil, got:\n%s", body)
	}
}

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
