//ff:func feature=stml-gen type=test control=sequence
//ff:what fnParam이 비어있지 않을 때 arrow function 형태인지 검증
package stml

import "testing"

func TestRenderMutationFnExpr_Arrow(t *testing.T) {
	got := renderMutationFnExpr("(data)", "UpdateRoom", "{ ...data, room_id: Number(RoomID) }")
	want := "(data) => api.UpdateRoom({ ...data, room_id: Number(RoomID) })"
	if got != want {
		t.Errorf("renderMutationFnExpr arrow = %q, want %q", got, want)
	}
}
