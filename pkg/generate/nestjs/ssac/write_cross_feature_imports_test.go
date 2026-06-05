//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteCrossFeatureImports — TestWriteCrossFeatureImports — cross-feature module import 문 출력 검증

package ssac

import (
	"strings"
	"testing"
)

func TestWriteCrossFeatureImports(t *testing.T) {
	var b strings.Builder
	writeCrossFeatureImports(&b, []string{"billing", "shipping"})

	out := b.String()
	want := "import { BillingModule } from '../billing/billing.module';\n" +
		"import { ShippingModule } from '../shipping/shipping.module';\n"
	if out != want {
		t.Errorf("got %q\nwant %q", out, want)
	}
}
