//ff:func feature=gen-react type=test control=sequence
//ff:what hasClaimsCaptures — claims 캡처 게이트 양성/토큰 캡처만/페이지 부재 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestHasClaimsCaptures(t *testing.T) {
	with := []stml.PageSpec{{Actions: []stml.ActionBlock{{
		Captures: []stml.CaptureBind{{RespField: "role", Sink: "auth.claims.role"}},
	}}}}
	without := []stml.PageSpec{{Actions: []stml.ActionBlock{{
		Captures: []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}},
	}}}}
	if !hasClaimsCaptures(with) {
		t.Errorf("claims capture not detected")
	}
	if hasClaimsCaptures(without) {
		t.Errorf("token capture must not count as a claims capture")
	}
	if hasClaimsCaptures(nil) {
		t.Errorf("no pages: want false")
	}
}
