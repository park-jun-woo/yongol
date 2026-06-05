//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteAuthResourceIDWithOwnership — TestWriteAuthResourceID — Ownership 유무에 따른 resourceId/owners 인자 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteAuthResourceIDWithOwnership(t *testing.T) {
	var b strings.Builder
	op := &ir.AuthOp{
		Resource: "note",
		Ownership: &ir.OwnershipInfo{
			ResourcePK:  "id",
			OwnerColumn: "owner_id",
		},
	}
	writeAuthResourceID(&b, op, "  ")

	out := b.String()
	if !strings.Contains(out, "resourceId: String(params.id),") {
		t.Errorf("missing resourceId line, got %q", out)
	}
	if !strings.Contains(out, "owners: { note: { owner_id: owner?.owner_id } },") {
		t.Errorf("missing owners line, got %q", out)
	}
}
