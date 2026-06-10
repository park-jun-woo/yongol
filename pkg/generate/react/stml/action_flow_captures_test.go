//ff:func feature=stml-gen type=test control=sequence
//ff:what actionFlowCaptures — bearer 모드에서만 캡처 반환·cookie 모드는 nil 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionFlowCaptures(t *testing.T) {
	binds := []stmlparser.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "refresh_token", Sink: "auth.refresh"},
	}
	a := stmlparser.ActionBlock{Captures: binds}

	// bearer mode -> returns the captures
	got := actionFlowCaptures(a, true)
	if len(got) != 2 || got[0] != binds[0] || got[1] != binds[1] {
		t.Errorf("bearer mode captures = %+v, want %+v", got, binds)
	}

	// cookie mode -> nil (captures ignored, TM-24 diagnoses at validate)
	if got := actionFlowCaptures(a, false); got != nil {
		t.Errorf("cookie mode captures = %+v, want nil", got)
	}
}
