//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-71 — s71BuildScope 단위 테스트 (scope 변수 누적 검증)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestS71BuildScope(t *testing.T) {
	t.Run("Always_has_request", func(t *testing.T) {
		fn := ssac.ServiceFunc{Name: "X"}
		scope := s71BuildScope(fn, 0)
		if !scope["request"] {
			t.Error("scope should always contain 'request'")
		}
	})
	t.Run("Subscribe_adds_message", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Name:      "OnOrder",
			Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"},
		}
		scope := s71BuildScope(fn, 0)
		if !scope["message"] {
			t.Error("scope should contain 'message' for subscribe funcs")
		}
	})
	t.Run("Auth_adds_currentUser", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Name: "X",
			Sequences: []ssac.Sequence{
				{Type: "auth", Line: 3},
				{Type: "get", Line: 5},
			},
		}
		scope := s71BuildScope(fn, 2)
		if !scope["currentUser"] {
			t.Error("scope should contain 'currentUser' after @auth")
		}
	})
	t.Run("Result_var_in_scope", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Name: "X",
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Result: &ssac.Result{Var: "course", Type: "Course"}},
				{Type: "empty", Line: 5, Target: "course"},
			},
		}
		scope := s71BuildScope(fn, 1)
		if !scope["course"] {
			t.Error("scope should contain 'course' after @get result binding")
		}
	})
	t.Run("Upto_boundary", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Name: "X",
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Result: &ssac.Result{Var: "a", Type: "A"}},
				{Type: "get", Line: 5, Result: &ssac.Result{Var: "b", Type: "B"}},
			},
		}
		scope := s71BuildScope(fn, 1)
		if !scope["a"] {
			t.Error("scope should contain 'a' at upto=1")
		}
		if scope["b"] {
			t.Error("scope should NOT contain 'b' at upto=1")
		}
	})
}
