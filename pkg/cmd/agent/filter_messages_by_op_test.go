//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestFilterMessagesByOp — op 를 언급하는 메시지만 필터링 검증

package agent

import "testing"

func TestFilterMessagesByOp(t *testing.T) {
	msgs := []string{"CreateUser bad", "ListOrders bad", "CreateUser worse"}
	got := filterMessagesByOp(msgs, "CreateUser")
	if len(got) != 2 {
		t.Errorf("filterMessagesByOp = %v, want 2 entries", got)
	}
	if got := filterMessagesByOp(msgs, "Nope"); len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}
