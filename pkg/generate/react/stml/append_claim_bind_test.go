//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what appendClaimBind — claims sink 만 추가, token/refresh/비식별자 sink 는 목록 불변 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAppendClaimBind(t *testing.T) {
	claim := stmlparser.CaptureBind{RespField: "role", Sink: "auth.claims.role"}
	got := appendClaimBind(nil, claim)
	if len(got) != 1 || got[0] != claim {
		t.Errorf("claims sink: got %+v, want [%+v]", got, claim)
	}

	// non-claims sinks leave the list untouched
	base := []stmlparser.CaptureBind{claim}
	for _, sink := range []string{"auth.token", "auth.refresh", "auth.claims.", "auth.claims.ro-le", "other"} {
		got := appendClaimBind(base, stmlparser.CaptureBind{RespField: "x", Sink: sink})
		if len(got) != 1 || got[0] != claim {
			t.Errorf("sink %q: got %+v, want base unchanged", sink, got)
		}
	}
}
