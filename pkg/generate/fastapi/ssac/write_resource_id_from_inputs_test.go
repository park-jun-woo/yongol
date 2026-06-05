//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteResourceIDFromInputs — writeResourceIDFromInputs ResourceID 입력→resource_id 인자 출력 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteResourceIDFromInputs(t *testing.T) {
	t.Run("Present", func(t *testing.T) {
		var b strings.Builder
		inputs := []ir.FieldArg{
			{Key: "role", Literal: "admin"},
			{Key: "ResourceID", Literal: "order_id"},
		}
		writeResourceIDFromInputs(&b, inputs, "  ")
		want := "      resource_id=str(order_id),\n"
		if got := b.String(); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Absent", func(t *testing.T) {
		var b strings.Builder
		writeResourceIDFromInputs(&b, []ir.FieldArg{{Key: "role", Literal: "admin"}}, "")
		if b.String() != "" {
			t.Errorf("expected empty when no ResourceID, got %q", b.String())
		}
	})
}
