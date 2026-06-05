//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteAuthInputs — writeAuthInputs authz_check 인자 출력·ResourceID 스킵 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteAuthInputs(t *testing.T) {
	var b strings.Builder
	inputs := []ir.FieldArg{
		{Key: "role", Literal: "admin", IsQuoted: true},
		{Key: "ResourceID", Literal: "order"},
	}
	writeAuthInputs(&b, inputs, "  ")
	got := b.String()
	want := "      role=\"admin\",\n"
	if got != want {
		t.Errorf("writeAuthInputs() = %q, want %q", got, want)
	}
	if strings.Contains(got, "resource_id") || strings.Contains(got, "ResourceID") {
		t.Errorf("ResourceID input should be skipped, got %q", got)
	}
}
