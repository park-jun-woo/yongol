//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWritePlanComponentImportsEmpty — TestWritePlanComponentImports — plan 별 Controller/Service import 문 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWritePlanComponentImportsEmpty(t *testing.T) {
	var b strings.Builder
	writePlanComponentImports(&b, nil)
	if b.String() != "" {
		t.Errorf("got %q, want empty", b.String())
	}
}
