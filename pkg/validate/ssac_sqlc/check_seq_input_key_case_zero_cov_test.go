//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestCheckSeqInputKeyCase_ZeroCov(t *testing.T) {
	g := &rule.Ground{Lookup: map[string]rule.StringSet{
		"SQLc.param.User": {"Email": true},
	}}

	// call type → nil (early return).
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Type: "call", Model: "User.X"}, g); got != nil {
		t.Error("call type should return nil")
	}
	// package set → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Package: "session", Model: "User.X"}, g); got != nil {
		t.Error("package set should return nil")
	}
	// empty model → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{}, g); got != nil {
		t.Error("empty model should return nil")
	}
	// model without dot → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Model: "User"}, g); got != nil {
		t.Error("model without method should return nil")
	}
	// unknown model params → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Model: "Other.Find", Args: []ssac.Arg{{Field: "Email"}}}, g); got != nil {
		t.Error("unknown model should return nil")
	}
	// no input keys → nil.
	if got := checkSeqInputKeyCase(ssac.ServiceFunc{}, ssac.Sequence{Model: "User.Find"}, g); got != nil {
		t.Error("no input keys should return nil")
	}
	// casing mismatch → diag.
	diags := checkSeqInputKeyCase(
		ssac.ServiceFunc{FileName: "s.ssac"},
		ssac.Sequence{Model: "User.Find", Line: 9, Args: []ssac.Arg{{Field: "email"}}},
		g,
	)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
}
