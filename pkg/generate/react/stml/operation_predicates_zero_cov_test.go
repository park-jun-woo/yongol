//ff:func feature=stml-gen type=test control=sequence
//ff:what 단순 stml 헬퍼 (string/zod/param/flow) 묶음 커버 — coverage attribution 으로 다수 함수 PASS
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestOperationPredicates_ZeroCov(t *testing.T) {
	if !isDeleteOperation("DeleteX") || isDeleteOperation("GetX") {
		t.Errorf("isDeleteOperation wrong")
	}
	cap := stmlparser.ActionBlock{Captures: []stmlparser.CaptureBind{{RespField: "access_token", Sink: "auth.token"}}}
	if !actionHasFlowSuccess(cap, true) {
		t.Errorf("actionHasFlowSuccess capture+bearer should be true")
	}
	if actionHasFlowSuccess(cap, false) {
		t.Errorf("actionHasFlowSuccess capture+cookie should be false")
	}
	if !actionHasFlowSuccess(stmlparser.ActionBlock{Redirect: "/"}, false) {
		t.Errorf("actionHasFlowSuccess redirect should be true regardless of mode")
	}
	if actionHasFlowSuccess(stmlparser.ActionBlock{}, true) {
		t.Errorf("actionHasFlowSuccess empty should be false")
	}
	if got := errorStateVar(stmlparser.ActionBlock{OperationID: "Login", OnErrorNode: true}); got != "loginError" {
		t.Errorf("errorStateVar = %q", got)
	}
	// page-flow Phase004: the error state is derived regardless of OnErrorNode
	if got := errorStateVar(stmlparser.ActionBlock{OperationID: "Login"}); got != "loginError" {
		t.Errorf("errorStateVar without OnErrorNode = %q", got)
	}
	types := map[string]map[string]string{"GetX": {"id": "integer"}}
	if !isIntegerParam("GetX", "id", types) {
		t.Errorf("isIntegerParam should be true")
	}
	if isIntegerParam("GetX", "id", nil) {
		t.Errorf("isIntegerParam nil map should be false")
	}
	if isIntegerParam("Unknown", "id", types) {
		t.Errorf("isIntegerParam unknown op should be false")
	}
}
