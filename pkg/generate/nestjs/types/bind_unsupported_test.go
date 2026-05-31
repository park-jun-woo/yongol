//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBindUnsupported — bindUnsupported 거절 바인딩 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindUnsupported_ZeroCov(t *testing.T) {
	b := bindUnsupported(typemap.FamilyUnsupported)
	if b.Supported {
		t.Errorf("bindUnsupported should not be Supported")
	}
}
