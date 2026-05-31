//ff:func feature=gen-gogin type=test control=sequence
//ff:what ownershipLookup 단위 테스트 (nil/zero-rid 조기반환 + UUID OwnerLookup 방출)
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestOwnershipLookup(t *testing.T) {
	newG := func() *methodGen {
		return &methodGen{
			FuncName:     "DeleteWorkflow",
			UseTx:        true,
			DeclaredVars: map[string]bool{},
			DDLTables: []ddl.Table{
				{
					Name: "workflows",
					Columns: map[string]ddl.Column{
						"id": {Name: "id", RawType: "UUID", NotNull: true},
					},
				},
			},
		}
	}

	t.Run("nil mapping → no lines, nil owners", func(t *testing.T) {
		g := newG()
		lines, owners, imps := g.ownershipLookup(ssacparser.Sequence{}, nil, 403, "Forbidden")
		if lines != nil || owners != "nil" || imps != nil {
			t.Errorf("got (%v,%q,%v), want (nil,nil,nil)", lines, owners, imps)
		}
	})

	t.Run("missing/zero ResourceID → nil owners", func(t *testing.T) {
		g := newG()
		mapping := &rego.OwnershipMapping{Resource: "workflow"}
		seq := ssacparser.Sequence{Inputs: map[string]string{"ResourceID": "0"}}
		lines, owners, _ := g.ownershipLookup(seq, mapping, 403, "Forbidden")
		if lines != nil || owners != "nil" {
			t.Errorf("zero rid should short-circuit, got (%v,%q)", lines, owners)
		}
	})

	t.Run("uuid resource emits OwnerLookup + owners map", func(t *testing.T) {
		g := newG()
		mapping := &rego.OwnershipMapping{Resource: "workflow"}
		seq := ssacparser.Sequence{Inputs: map[string]string{"ResourceID": "request.id"}}
		lines, owners, imps := g.ownershipLookup(seq, mapping, 403, "Forbidden")
		joined := strings.Join(lines, "\n")
		if !strings.Contains(joined, "qtx.OwnerLookupWorkflow(ctx,") {
			t.Errorf("missing OwnerLookup call:\n%s", joined)
		}
		if !strings.Contains(joined, "DeleteWorkflow403JSONResponse") {
			t.Errorf("missing error return with status:\n%s", joined)
		}
		if !strings.Contains(owners, "pgtypex.UUIDToString") {
			t.Errorf("uuid owner key should use pgtypex.UUIDToString, got %q", owners)
		}
		foundPgtypex := false
		for _, im := range imps {
			if strings.Contains(im, "pgtypex") {
				foundPgtypex = true
			}
		}
		if !foundPgtypex {
			t.Errorf("expected pgtypex import, got %v", imps)
		}
	})
}
