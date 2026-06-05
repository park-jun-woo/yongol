//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteAuthInputsOnlyResourceID — TestWriteAuthInputs — authz.check 인자 줄 출력(ResourceID 제외) 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteAuthInputsOnlyResourceID(t *testing.T) {
	var b strings.Builder
	writeAuthInputs(&b, []ir.FieldArg{{Key: "ResourceID"}}, "  ")
	if b.String() != "" {
		t.Errorf("got %q, want empty when only ResourceID", b.String())
	}
}
