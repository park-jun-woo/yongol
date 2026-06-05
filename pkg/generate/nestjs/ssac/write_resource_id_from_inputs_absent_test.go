//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteResourceIDFromInputsAbsent — TestWriteResourceIDFromInputs — Inputs 중 ResourceID 항목을 resourceId 인자로 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteResourceIDFromInputsAbsent(t *testing.T) {
	var b strings.Builder
	writeResourceIDFromInputs(&b, []ir.FieldArg{{Key: "OrgId"}}, "  ")
	if b.String() != "" {
		t.Errorf("got %q, want empty when no ResourceID input", b.String())
	}
}
