//ff:func feature=stml-gen type=test control=sequence
//ff:what fnParam이 빈 문자열(body only)일 때 api.X 직접 참조 형태인지 검증
package stml

import "testing"

func TestRenderMutationFnExpr_BodyOnly(t *testing.T) {
	got := renderMutationFnExpr("", "CreateRoom", "")
	want := "api.CreateRoom"
	if got != want {
		t.Errorf("renderMutationFnExpr body only = %q, want %q", got, want)
	}
}
