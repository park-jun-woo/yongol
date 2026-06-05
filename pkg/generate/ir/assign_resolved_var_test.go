//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAssignResolvedVar -- 변수명 충돌 해소·rename 등록·빈 이름/무충돌 분기 검증
package ir

import "testing"

func TestAssignResolvedVar(t *testing.T) {
	t.Run("EmptyNoop", func(t *testing.T) {
		name := ""
		renames := map[string]string{}
		assignResolvedVar(&name, map[string]bool{}, renames)
		if name != "" || len(renames) != 0 {
			t.Errorf("empty name should be a noop, got %q renames %v", name, renames)
		}
	})

	t.Run("NoCollision", func(t *testing.T) {
		name := "order"
		declared := map[string]bool{}
		renames := map[string]string{}
		assignResolvedVar(&name, declared, renames)
		if name != "order" {
			t.Errorf("name = %q, want order", name)
		}
		if len(renames) != 0 {
			t.Errorf("no rename expected, got %v", renames)
		}
		if !declared["order"] {
			t.Error("order should be registered as declared")
		}
	})

	t.Run("CollisionRenamed", func(t *testing.T) {
		name := "order"
		declared := map[string]bool{"order": true}
		renames := map[string]string{}
		assignResolvedVar(&name, declared, renames)
		if name != "order_result" {
			t.Errorf("name = %q, want order_result", name)
		}
		if renames["order"] != "order_result" {
			t.Errorf("renames = %v, want order->order_result", renames)
		}
	})
}
