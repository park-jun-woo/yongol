//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBuildPrismaQueryOptsWithWhereAndPagination — TestBuildPrismaQueryOpts — GetOp → where + pagination 옵션 목록 구성 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestBuildPrismaQueryOptsWithWhereAndPagination(t *testing.T) {
	op := &ir.GetOp{
		Args: []ir.FieldArg{
			{Key: "OwnerId", ColumnName: "owner_id", Location: ir.LocUser},
		},
		PaginationArgs: []ir.FieldArg{
			{Key: "Limit", ColumnName: "limit", Location: ir.LocQuery},
		},
	}

	opts := buildPrismaQueryOpts(op)

	if len(opts) != 2 {
		t.Fatalf("opts len = %d, want 2 (%v)", len(opts), opts)
	}
	if opts[0] != "where: { owner_id: user.owner_id }" {
		t.Errorf("opts[0] = %q", opts[0])
	}
	if opts[1] != "take: query.limit" {
		t.Errorf("opts[1] = %q", opts[1])
	}
}
