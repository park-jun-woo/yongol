//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteResourceIDFromInputsPresent — TestWriteResourceIDFromInputs — Inputs 중 ResourceID 항목을 resourceId 인자로 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteResourceIDFromInputsPresent(t *testing.T) {
	var b strings.Builder
	inputs := []ir.FieldArg{
		{Key: "OrgId", Location: ir.LocUser, ColumnName: "org_id"},
		{Key: "ResourceID", Location: ir.LocPath, ColumnName: "id"},
	}
	writeResourceIDFromInputs(&b, inputs, "  ")

	want := "    resourceId: String(params.id),\n"
	if b.String() != want {
		t.Errorf("got %q, want %q", b.String(), want)
	}
}
