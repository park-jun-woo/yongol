//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteAuthResourceIDNoOwnership — TestWriteAuthResourceID — Ownership 유무에 따른 resourceId/owners 인자 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteAuthResourceIDNoOwnership(t *testing.T) {
	var b strings.Builder
	op := &ir.AuthOp{
		Resource: "note",
		Inputs:   []ir.FieldArg{{Key: "ResourceID", Location: ir.LocPath, ColumnName: "id"}},
	}
	writeAuthResourceID(&b, op, "  ")

	out := b.String()
	if !strings.Contains(out, "resourceId: String(params.id),") {
		t.Errorf("expected resourceId from inputs, got %q", out)
	}
	if strings.Contains(out, "owners:") {
		t.Errorf("should not write owners when no ownership, got %q", out)
	}
}
