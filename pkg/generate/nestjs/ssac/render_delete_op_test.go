//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderDeleteOp(t *testing.T) {
	var b strings.Builder
	renderDeleteOp(&b, nil, "  ", "this.prisma")
	if b.String() != "" {
		t.Errorf("nil op")
	}
	b.Reset()
	renderDeleteOp(&b, &ir.DeleteOp{Model: "Course", Args: []ir.FieldArg{{Key: "id", IsPK: true}}}, "  ", "this.prisma")
	if !strings.Contains(b.String(), "this.prisma.course.delete(") {
		t.Errorf("pk delete = %q", b.String())
	}
	b.Reset()
	renderDeleteOp(&b, &ir.DeleteOp{Model: "Course", Args: []ir.FieldArg{{Key: "slug"}}}, "  ", "this.prisma")
	if !strings.Contains(b.String(), "deleteMany(") {
		t.Errorf("non-pk deleteMany = %q", b.String())
	}
}
