//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestSplitPutParts — PutOp Args 를 IsPK 기준 where/data Prisma 조각으로 분리 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestSplitPutParts(t *testing.T) {
	args := []ir.FieldArg{
		{Key: "Id", ColumnName: "id", IsPK: true, Location: ir.LocPath},
		{Key: "Title", ColumnName: "title", Location: ir.LocBody},
	}

	whereParts, dataParts := splitPutParts(args)

	if len(whereParts) != 1 || whereParts[0] != "id: params.id" {
		t.Errorf("whereParts = %v, want [id: params.id]", whereParts)
	}
	if len(dataParts) != 1 || dataParts[0] != "title: body.title" {
		t.Errorf("dataParts = %v, want [title: body.title]", dataParts)
	}
}
