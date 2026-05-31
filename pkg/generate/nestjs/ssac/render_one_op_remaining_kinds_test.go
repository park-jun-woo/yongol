//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderOneOp_RemainingKinds(t *testing.T) {
	// Dispatch coverage for the remaining renderOneOp branches.
	ops := []ir.Op{
		{Kind: ir.OpGet, Get: &ir.GetOp{VarName: "c", VarType: "Course", Model: "Course"}},
		{Kind: ir.OpPost, Post: &ir.PostOp{VarName: "c", Model: "Course"}},
		{Kind: ir.OpPut, Put: &ir.PutOp{Model: "Course", Args: []ir.FieldArg{{Key: "id", IsPK: true}}}},
		{Kind: ir.OpAuth, Auth: &ir.AuthOp{Action: "read", Resource: "course", Message: "denied", StatusCode: 403}},
		{Kind: ir.OpState, State: &ir.StateOp{Diagram: "d", Transition: "go", Message: "bad", StatusCode: 409}},
		{Kind: ir.OpVerifyPassword, VerifyPW: &ir.VerifyPasswordOp{Model: "User", EmailCol: "Email", EmailExpr: "request.email", HashCol: "PasswordHash", PasswordExpr: "request.password", ResultVar: "u", Message: "no"}},
		{Kind: ir.OpResponse, Response: &ir.ResponseOp{SingleVar: "c"}},
		{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Model: "Course", Args: []ir.FieldArg{{Key: "id", IsPK: true}}}},
	}
	for _, op := range ops {
		var b strings.Builder
		// must not panic; each branch dispatches to its renderer
		renderOneOp(&b, op, "  ", "this.prisma")
	}
}
