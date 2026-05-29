//ff:func feature=stml-gen type=test control=sequence
//ff:what void(body 없음) mutationFn이 arrow function 형태인지 검증
package stml

import "testing"

func TestRenderMutationFnExpr_Void(t *testing.T) {
	got := renderMutationFnExpr("()", "LogoutSession", "")
	want := "() => api.LogoutSession()"
	if got != want {
		t.Errorf("renderMutationFnExpr void = %q, want %q", got, want)
	}
}
