//ff:func feature=gen-react type=test control=sequence
//ff:what actionHasRefreshCapture — auth.refresh sink 보유/미보유 액션 판정 검증 (BUG-135)
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionHasRefreshCapture(t *testing.T) {
	with := stml.ActionBlock{Captures: []stml.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "refresh_token", Sink: "auth.refresh"},
	}}
	if !actionHasRefreshCapture(with) {
		t.Errorf("expected true for an action declaring an auth.refresh capture")
	}
	without := stml.ActionBlock{Captures: []stml.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
	}}
	if actionHasRefreshCapture(without) {
		t.Errorf("expected false for an action with only an auth.token capture")
	}
}
