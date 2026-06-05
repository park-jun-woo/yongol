//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteAuthInputs — TestWriteAuthInputs — authz.check 인자 줄 출력(ResourceID 제외) 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteAuthInputs(t *testing.T) {
	var b strings.Builder
	inputs := []ir.FieldArg{
		{Key: "ResourceID", Location: ir.LocPath}, // skipped
		{Key: "OrgId", ColumnName: "org_id", Location: ir.LocUser},
	}
	writeAuthInputs(&b, inputs, "  ")

	out := b.String()
	if strings.Contains(out, "resource_id") || strings.Contains(out, "ResourceID") {
		t.Errorf("ResourceID should be skipped, got %q", out)
	}
	want := "    org_id: user.org_id,\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}
