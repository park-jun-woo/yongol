//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBuildWhereParts — TestBuildWhereParts — FieldArg 목록 → Prisma where "key: value" 조각 생성 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestBuildWhereParts(t *testing.T) {
	args := []ir.FieldArg{
		{Key: "OwnerId", ColumnName: "owner_id", Location: ir.LocUser},
		{Key: "Status", ColumnName: "status", Location: ir.LocQuery},
	}

	parts := buildWhereParts(args)

	if len(parts) != 2 {
		t.Fatalf("parts len = %d, want 2 (%v)", len(parts), parts)
	}
	if parts[0] != "owner_id: user.owner_id" {
		t.Errorf("parts[0] = %q, want owner_id: user.owner_id", parts[0])
	}
	if parts[1] != "status: query.status" {
		t.Errorf("parts[1] = %q, want status: query.status", parts[1])
	}
}
