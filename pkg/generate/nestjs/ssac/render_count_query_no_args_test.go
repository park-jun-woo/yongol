//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderCountQueryNoArgs — TestRenderCountQuery — GetOp(IsCount) → Prisma count({ where }) / count() 렌더 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderCountQueryNoArgs(t *testing.T) {
	var b strings.Builder
	op := &ir.GetOp{VarName: "total"}
	renderCountQuery(&b, op, "  ", "this.prisma", "course")

	want := "  const total = await this.prisma.course.count();\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
