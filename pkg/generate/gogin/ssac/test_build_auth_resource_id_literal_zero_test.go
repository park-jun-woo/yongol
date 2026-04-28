//ff:func feature=gen-gogin type=test control=sequence topic=authz
//ff:what TestBuildAuth_ResourceIDLiteralZero_SkipsOwnerLookup — ResourceID==0 리터럴 skip (BUG-033)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
