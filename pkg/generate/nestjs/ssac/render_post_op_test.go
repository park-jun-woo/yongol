//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderPostOp(t *testing.T) {
	var b strings.Builder
	renderPostOp(&b, nil, "  ", "this.prisma")
	if b.String() != "" {
		t.Errorf("nil post")
	}
	b.Reset()
	renderPostOp(&b, &ir.PostOp{VarName: "course", Model: "Course", Args: []ir.FieldArg{{Key: "title", ColumnName: "title", Source: "body.title"}}}, "  ", "this.prisma")
	if !strings.Contains(b.String(), "const course = await this.prisma.course.create(") {
		t.Errorf("post op = %q", b.String())
	}
}
