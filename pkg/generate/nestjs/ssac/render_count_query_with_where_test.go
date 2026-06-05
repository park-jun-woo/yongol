//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderCountQueryWithWhere — TestRenderCountQuery — GetOp(IsCount) → Prisma count({ where }) / count() 렌더 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderCountQueryWithWhere(t *testing.T) {
	var b strings.Builder
	op := &ir.GetOp{
		VarName: "total",
		Args:    []ir.FieldArg{{Key: "OwnerId", ColumnName: "owner_id", Location: ir.LocUser}},
	}
	renderCountQuery(&b, op, "  ", "this.prisma", "course")

	want := "  const total = await this.prisma.course.count({ where: { owner_id: user.owner_id } });\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
