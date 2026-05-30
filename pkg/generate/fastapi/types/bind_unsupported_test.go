//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestBindUnsupported — FamilyUnsupported → Supported=false 바인딩

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

func TestBindUnsupported(t *testing.T) {
	b := bindUnsupported(typemap.FamilyUnsupported)
	if b.Supported {
		t.Errorf("Supported = true, want false")
	}
	if b.Family != typemap.FamilyUnsupported {
		t.Errorf("Family = %v, want FamilyUnsupported", b.Family)
	}
}
