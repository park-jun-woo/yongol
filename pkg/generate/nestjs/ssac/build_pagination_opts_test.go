//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBuildPaginationOpts — PaginationArgs → Prisma take/skip/cursor 옵션 목록 생성 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestBuildPaginationOpts(t *testing.T) {
	pagArgs := []ir.FieldArg{
		{Key: "Limit", ColumnName: "limit", Location: ir.LocQuery},
		{Key: "Cursor", ColumnName: "cursor", Location: ir.LocQuery},
	}

	opts := buildPaginationOpts(pagArgs)

	if len(opts) != 2 {
		t.Fatalf("opts len = %d, want 2 (%v)", len(opts), opts)
	}
	if opts[0] != "take: query.limit" {
		t.Errorf("opts[0] = %q, want take: query.limit", opts[0])
	}
	want := "cursor: query.cursor ? { id: query.cursor } : undefined"
	if opts[1] != want {
		t.Errorf("opts[1] = %q, want %q", opts[1], want)
	}
}
