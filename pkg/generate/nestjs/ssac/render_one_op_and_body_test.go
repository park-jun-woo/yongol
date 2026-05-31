//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderOneOpAndBody(t *testing.T) {
	var b strings.Builder
	ops := []ir.Op{
		{Kind: ir.OpEmpty, Empty: &ir.EmptyOp{VarName: "c", Message: "nf", StatusCode: 404}},
		{Kind: ir.OpExists, Exists: &ir.ExistsOp{VarName: "d", Message: "cf", StatusCode: 409}},
		{Kind: ir.OpCall, Call: &ir.CallOp{Function: "Run"}},
		{Kind: ir.OpEval, Eval: &ir.EvalOp{Function: "Check", Message: "bad", StatusCode: 400}},
		{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "t"}},
	}
	renderOpsBody(&b, ops, "  ", "this.prisma")
	out := b.String()
	for _, want := range []string{"if (!c)", "if (d)", "await run()", "await check()", "this.queue.publish('t'"} {
		if !strings.Contains(out, want) {
			t.Errorf("renderOpsBody missing %q in %q", want, out)
		}
	}
}
