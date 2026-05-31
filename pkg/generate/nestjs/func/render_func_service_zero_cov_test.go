//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderFuncModule_ZeroCov — 외부 패키지 stub module 소스 생성
package funcstub

import (
	"strings"
	"testing"
)

func TestRenderFuncService_ZeroCov(t *testing.T) {
	out := RenderFuncService("billing", []string{"Charge", "Refund"})
	for _, want := range []string{
		"export class BillingService",
		"async charge(",
		"async refund(",
		"not implemented",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderFuncService missing %q\n%s", want, out)
		}
	}
	// No methods → empty body
	empty := RenderFuncService("misc", nil)
	if !strings.Contains(empty, "export class MiscService") {
		t.Errorf("RenderFuncService empty missing class:\n%s", empty)
	}
}
