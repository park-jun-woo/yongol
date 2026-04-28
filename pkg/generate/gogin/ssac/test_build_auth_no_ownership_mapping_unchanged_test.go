//ff:func feature=gen-gogin type=test control=sequence topic=authz
//ff:what TestBuildAuth_NoOwnershipMappingUnchanged — 매핑 없으면 OwnerLookup 미생성

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
