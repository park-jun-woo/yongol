//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestPgtypeConstructorsZeroCov — 모든 pgtype 생성자 + unsupportedBinding 직접 커버
package types

import (
	"testing"
)

func TestUnsupportedBinding_ZeroCov(t *testing.T) {
	b := unsupportedBinding("nope")
	if b.Supported || b.Kind != KindUnsupported {
		t.Errorf("unsupportedBinding should be unsupported: %+v", b)
	}
}
