//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXsd55DDLToModelRefPluralTablesStillMatched — plural tables still matched 서브테스트
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func subtestTestXsd55DDLToModelRefPluralTablesStillMatched(t *testing.T) {

	fs := newXsd55Fullstack(
		rule.StringSet{},
		ddl.Table{Name: "users"},
		ddl.Table{Name: "bid_requests"},
		ddl.Table{Name: "address"},   // ss preserved
		ddl.Table{Name: "companies"}, // ies → y
	)
	fs.ServiceFuncs = []ssac.ServiceFunc{
		{Sequences: []ssac.Sequence{{Model: "User.Get"}}},
		{Sequences: []ssac.Sequence{{Model: "BidRequest.List"}}},
		{Sequences: []ssac.Sequence{{Result: &ssac.Result{Type: "Address"}}}},
		{Sequences: []ssac.Sequence{{Result: &ssac.Result{Type: "Company"}}}},
	}
	if diags := xsd55DDLToModelRef(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags for matched plural tables, got %d: %v", len(diags), diags)
	}

}
