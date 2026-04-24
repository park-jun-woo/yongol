//ff:func feature=gen-gogin type=test control=sequence topic=authz
//ff:what TestBuildAuth — owner lookup 은 := 로 선언해야 함 (BUG-029 regression)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildAuth_OwnerLookupDeclaresWithWalrus verifies that when an
// `@ownership` mapping applies to the @auth sequence, buildAuth emits
// the OwnerLookup<Resource> call with Go's short-declaration operator
// `:=` — even when the enclosing handler already declared err earlier
// (UseTx=true / FirstErr=false). The owner variable is introduced by
// this line, so `=` would yield `undefined: owner<Resource>`. (BUG-029)
func TestBuildAuth_OwnerLookupDeclaresWithWalrus(t *testing.T) {
	g := &methodGen{
		FuncName:   "CreateWorkflow",
		FileName:   "workflow_service.ssac",
		ModulePath: "example.com/zenflow",
		UseTx:      true,
		FirstErr:   false, // tx preamble already used :=
		Ownerships: []rego.OwnershipMapping{
			{Resource: "workflow", Table: "workflows", Column: "org_id"},
		},
	}
	seq := ssacparser.Sequence{
		Type:     "auth",
		Action:   "CreateWorkflow",
		Resource: "workflow",
		Inputs: map[string]string{
			"ResourceID": "body.WorkflowID",
		},
	}
	lines, _ := g.buildAuth(seq)
	body := strings.Join(lines, "\n")

	if !strings.Contains(body, "ownerWorkflow, err := qtx.OwnerLookupWorkflow(ctx,") {
		t.Fatalf("expected owner lookup declared with := on qtx, got:\n%s", body)
	}
	// The subsequent authz.Check must reuse err with `=` (err was
	// already declared in the owner-lookup line above).
	if !strings.Contains(body, "_, err = authz.Check(") {
		t.Fatalf("expected authz.Check to reassign err with =, got:\n%s", body)
	}
}

// TestBuildAuth_NoOwnershipMappingUnchanged guards the non-ownership
// path: without a matching @ownership mapping no OwnerLookup line is
// emitted and Owners is passed as nil.
func TestBuildAuth_NoOwnershipMappingUnchanged(t *testing.T) {
	g := &methodGen{
		FuncName:   "ListWorkflows",
		FileName:   "workflow_service.ssac",
		ModulePath: "example.com/zenflow",
		UseTx:      false,
		FirstErr:   true,
		Ownerships: nil,
	}
	seq := ssacparser.Sequence{
		Type:     "auth",
		Action:   "ListWorkflows",
		Resource: "workflow",
	}
	lines, _ := g.buildAuth(seq)
	body := strings.Join(lines, "\n")

	if strings.Contains(body, "OwnerLookup") {
		t.Fatalf("no ownership mapping: must not emit OwnerLookup line, got:\n%s", body)
	}
	if !strings.Contains(body, "Owners: nil") {
		t.Fatalf("no ownership mapping: Owners must be nil, got:\n%s", body)
	}
}
