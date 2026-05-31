//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertPublish_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "CompleteOrder", []ssac.Sequence{
		{Type: ssac.SeqPost, Model: "Order.Complete",
			Inputs: map[string]string{"ID": "request.id"},
			Result: &ssac.Result{Var: "order", Type: "Order"}},
		{Type: ssac.SeqPublish, Topic: "order.completed",
			Inputs: map[string]string{"OrderID": "order.ID", "UserID": "currentUser.ID"}},
	})
	found := false
	for _, op := range plan.Ops {
		if op.Kind == OpPublish {
			found = true
		}
	}
	if !found {
		t.Errorf("expected an OpPublish in plan")
	}
}
