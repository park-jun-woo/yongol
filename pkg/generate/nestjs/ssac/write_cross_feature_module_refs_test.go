//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteCrossFeatureModuleRefs — TestWriteCrossFeatureModuleRefs — @Module imports 배열 cross-feature Module 참조 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWriteCrossFeatureModuleRefs(t *testing.T) {
	var b strings.Builder
	writeCrossFeatureModuleRefs(&b, []string{"billing", "shipping"})

	want := "    BillingModule,\n    ShippingModule,\n"
	if b.String() != want {
		t.Errorf("got %q\nwant %q", b.String(), want)
	}
}
