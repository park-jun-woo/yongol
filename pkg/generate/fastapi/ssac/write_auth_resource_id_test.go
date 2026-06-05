//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteAuthResourceID — writeAuthResourceID ownership 유무에 따른 resource_id/owners 출력 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteAuthResourceID(t *testing.T) {
	t.Run("WithOwnership", func(t *testing.T) {
		var b strings.Builder
		op := &ir.AuthOp{
			Resource: "project",
			Ownership: &ir.OwnershipInfo{
				ResourcePK:  "project_id",
				OwnerColumn: "owner_id",
			},
		}
		writeAuthResourceID(&b, op, "  ")
		got := b.String()
		want := "      resource_id=str(project_id),\n" +
			"      owners={\"project\": {\"owner_id\": owner}},\n"
		if got != want {
			t.Errorf("with ownership:\n got %q\nwant %q", got, want)
		}
	})

	t.Run("WithoutOwnership", func(t *testing.T) {
		var b strings.Builder
		op := &ir.AuthOp{
			Inputs: []ir.FieldArg{{Key: "ResourceID", Literal: "order_id"}},
		}
		writeAuthResourceID(&b, op, "")
		got := b.String()
		want := "    resource_id=str(order_id),\n"
		if got != want {
			t.Errorf("without ownership: got %q, want %q", got, want)
		}
	})
}
