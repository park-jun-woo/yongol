//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCheckSingleInputKeyCase_ZeroCov(t *testing.T) {
	fn := ssac.ServiceFunc{FileName: "svc.ssac"}
	seq := ssac.Sequence{Line: 5}

	// exact match → no diag.
	if _, fired := checkSingleInputKeyCase(fn, seq, "Email", map[string]bool{"Email": true}); fired {
		t.Error("exact match should not fire")
	}
	// snake match → no diag.
	if _, fired := checkSingleInputKeyCase(fn, seq, "BidAmount", map[string]bool{"bid_amount": true}); fired {
		t.Error("snake-equivalent should not fire")
	}
	// no match at all → no diag.
	if _, fired := checkSingleInputKeyCase(fn, seq, "Email", map[string]bool{"other": true}); fired {
		t.Error("no match should not fire")
	}
	// case-insensitive only match → fires.
	d, fired := checkSingleInputKeyCase(fn, seq, "EMAIL", map[string]bool{"Email": true})
	if !fired {
		t.Fatal("case-insensitive mismatch should fire")
	}
	if d.File != "svc.ssac" || d.Line != 5 {
		t.Errorf("diag fields = %+v", d)
	}
}
