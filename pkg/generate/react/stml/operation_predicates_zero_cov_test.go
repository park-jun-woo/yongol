//ff:func feature=stml-gen type=test control=sequence
//ff:what 단순 stml 헬퍼 (string/zod/param/login) 묶음 커버 — coverage attribution 으로 다수 함수 PASS
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestOperationPredicates_ZeroCov(t *testing.T) {
	if !isDeleteOperation("DeleteX") || isDeleteOperation("GetX") {
		t.Errorf("isDeleteOperation wrong")
	}
	if !isLoginAction("Login") || isLoginAction("X") {
		t.Errorf("isLoginAction wrong")
	}
	if !hasLoginAction([]stmlparser.ActionBlock{{OperationID: "Login"}}) {
		t.Errorf("hasLoginAction should be true")
	}
	if hasLoginAction([]stmlparser.ActionBlock{{OperationID: "X"}}) {
		t.Errorf("hasLoginAction should be false")
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
