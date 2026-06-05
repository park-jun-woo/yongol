//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteCrossFeatureModuleRefsEmpty — TestWriteCrossFeatureModuleRefs — @Module imports 배열 cross-feature Module 참조 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWriteCrossFeatureModuleRefsEmpty(t *testing.T) {
	var b strings.Builder
	writeCrossFeatureModuleRefs(&b, nil)
	if b.String() != "" {
		t.Errorf("got %q, want empty", b.String())
	}
}
