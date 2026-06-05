//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteCrossFeatureImportsEmpty — TestWriteCrossFeatureImports — cross-feature module import 문 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWriteCrossFeatureImportsEmpty(t *testing.T) {
	var b strings.Builder
	writeCrossFeatureImports(&b, nil)
	if b.String() != "" {
		t.Errorf("got %q, want empty", b.String())
	}
}
