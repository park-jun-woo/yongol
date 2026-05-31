//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderPublishOp(t *testing.T) {
	var b strings.Builder
	renderPublishOp(&b, nil, "  ")
	if b.String() != "" {
		t.Errorf("nil publish")
	}
	b.Reset()
	renderPublishOp(&b, &ir.PublishOp{Topic: "order.completed", Payload: []ir.FieldArg{{Key: "id", Literal: "1"}}}, "  ")
	out := b.String()
	if !strings.Contains(out, "this.queue.publish('order.completed'") || !strings.Contains(out, "id: 1") {
		t.Errorf("publish op = %q", out)
	}
}
