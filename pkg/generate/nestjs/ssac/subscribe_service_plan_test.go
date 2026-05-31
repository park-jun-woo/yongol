//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버
package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subscribeServicePlan() *ir.ServicePlan {
	return &ir.ServicePlan{
		OperationID: "OnCourseCreated",
		TriggerKind: ir.TriggerSubscribe,
		Topic:       "course.created",
		Feature:     "course",
		Ops: []ir.Op{
			{Kind: ir.OpPost, Post: &ir.PostOp{VarName: "log", Model: "Log"}},
		},
	}
}
