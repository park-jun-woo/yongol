//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func bnPlan() *ir.ServicePlan {
	return &ir.ServicePlan{
		OperationID: "CreateOrder",
		HTTPMethod:  "POST",
		TriggerKind: ir.TriggerHTTP,
		URLPath:     "/orders/:id/items",
		BodyFields:  []ir.BodyFieldMeta{{Name: "qty"}},
		PathParams:  []string{"id"},
		QueryParams: []ir.QueryParamMeta{{Name: "page", Type: "integer"}},
		Ops: []ir.Op{
			{Kind: ir.OpAuth, Auth: &ir.AuthOp{}},
			{Kind: ir.OpPublish, Publish: &ir.PublishOp{}},
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing"}},
		},
	}
}
