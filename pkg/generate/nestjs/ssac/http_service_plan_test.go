//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov2 — RenderService/RenderController 진입점으로 write* 함수 0% 커버
package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func httpServicePlan() *ir.ServicePlan {
	return &ir.ServicePlan{
		OperationID:     "CreateCourse",
		TriggerKind:     ir.TriggerHTTP,
		HTTPMethod:      "POST",
		URLPath:         "/courses",
		Feature:         "course",
		UsesTransaction: true,
		PathParams:      []string{"id"},
		QueryParams:     []ir.QueryParamMeta{{Name: "limit"}},
		BodyFields:      []ir.BodyFieldMeta{{Name: "title"}},
		Ops: []ir.Op{
			{Kind: ir.OpPost, Post: &ir.PostOp{VarName: "course", Model: "Course"}},
			{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "course.created"}},
			{Kind: ir.OpAuth, Auth: &ir.AuthOp{Action: "create", Resource: "course", Message: "denied", StatusCode: 403}},
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}},
		},
	}
}
