//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderFuncModule_ZeroCov — 외부 패키지 stub module 소스 생성
package funcstub

import (
	"strings"
	"testing"
)

func TestRenderFuncModule_ZeroCov(t *testing.T) {
	out := RenderFuncModule("billing", []string{"Charge"})
	for _, want := range []string{"BillingService", "export class BillingModule", "./billing.service", "@Global()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderFuncModule missing %q\n%s", want, out)
		}
	}
}
