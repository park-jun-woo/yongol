//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-37 — FK reference @get 에 @empty guard 필수 검증 (guard 있으면 pass, 없으면 warning)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS37FKReferenceGuard(t *testing.T) {
	t.Run("Fires_FK_ref_without_guard", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "User.FindByID", Result: &ssac.Result{Type: "User", Var: "user"}},
			{Type: "get", Line: 5, Model: "Order.FindByUserID", Args: []ssac.Arg{{Source: "user", Field: "ID"}}, Result: &ssac.Result{Type: "Order", Var: "order"}},
			{Type: "response", Line: 7},
		}}}}
		assertDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Passes_with_empty_guard", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "User.FindByID", Result: &ssac.Result{Type: "User", Var: "user"}},
			{Type: "get", Line: 5, Model: "Order.FindByUserID", Args: []ssac.Arg{{Source: "user", Field: "ID"}}, Result: &ssac.Result{Type: "Order", Var: "order"}},
			{Type: "empty", Line: 7, Target: "order"},
		}}}}
		assertNoDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Passes_with_exists_guard", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "User.FindByID", Result: &ssac.Result{Type: "User", Var: "user"}},
			{Type: "get", Line: 5, Model: "Order.FindByUserID", Args: []ssac.Arg{{Source: "user", Field: "ID"}}, Result: &ssac.Result{Type: "Order", Var: "order"}},
			{Type: "exists", Line: 7, Target: "order"},
		}}}}
		assertNoDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Passes_same_model", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "Order.FindByID", Result: &ssac.Result{Type: "Order", Var: "order"}},
			{Type: "get", Line: 5, Model: "Order.FindByCode", Args: []ssac.Arg{{Source: "order", Field: "Code"}}, Result: &ssac.Result{Type: "Order", Var: "order2"}},
		}}}}
		assertNoDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Passes_implicit", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "request", Field: "ID"}}, Result: &ssac.Result{Type: "Order", Var: "order"}},
		}}}}
		assertNoDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Skips_list", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "User.FindByID", Result: &ssac.Result{Type: "User", Var: "user"}},
			{Type: "get", Line: 5, Model: "Order.List", Args: []ssac.Arg{{Source: "user", Field: "ID"}}, Result: &ssac.Result{Type: "[]Order", Var: "orders"}},
		}}}}
		assertNoDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Skips_primitive", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "User.FindByID", Result: &ssac.Result{Type: "User", Var: "user"}},
			{Type: "get", Line: 5, Model: "Order.Count", Args: []ssac.Arg{{Source: "user", Field: "ID"}}, Result: &ssac.Result{Type: "int64", Var: "count"}},
		}}}}
		assertNoDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Fires_via_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "order.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "User.FindByID", Result: &ssac.Result{Type: "User", Var: "user"}},
			{Type: "get", Line: 5, Model: "Order.Find", Inputs: map[string]string{"userId": "user.ID"}, Result: &ssac.Result{Type: "Order", Var: "order"}},
		}}}}
		assertDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Fires_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "worker.ssac",
			Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "message", Field: "OrderID"}}, Result: &ssac.Result{Type: "Order", Var: "order"}},
				{Type: "get", Line: 5, Model: "User.FindByID", Args: []ssac.Arg{{Source: "order", Field: "UserID"}}, Result: &ssac.Result{Type: "User", Var: "user"}},
			}}}}
		assertDiag(t, s37FKReferenceGuard(fs), "[S-37]")
	})
	t.Run("Empty", func(t *testing.T) {
		assertNoDiag(t, s37FKReferenceGuard(&yongol.Fullstack{}), "[S-37]")
	})
}
