//ff:func feature=rule type=test control=sequence
//ff:what populateDDLFuncManaged — @func-managed 테이블 플래그가 Flags에 투영, 그 외엔 미설정
package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDLFuncManaged_Flag covers projection of FuncManaged onto
// Flags["funcManaged.<table>"] and confirms it does not leak into the
// archived/sensitive flag namespaces (XSD-55 한정).
func TestPopulateDDLFuncManaged_Flag(t *testing.T) {
	managed := ddl.Table{Name: "bids", FuncManaged: true}
	plain := ddl.Table{Name: "courses"}
	fs := newMinimalFullstack(withDDLTables(managed, plain))
	g := newGround()

	populateDDLFuncManaged(g, fs)

	if !g.Flags["funcManaged.bids"] {
		t.Errorf("funcManaged.bids flag missing")
	}
	if g.Flags["funcManaged.courses"] {
		t.Errorf("funcManaged.courses set unexpectedly")
	}
	// func-managed must not bleed into archived/sensitive flag namespaces.
	if g.Flags["archived.bids"] {
		t.Errorf("archived.bids set unexpectedly by func-managed")
	}
}
