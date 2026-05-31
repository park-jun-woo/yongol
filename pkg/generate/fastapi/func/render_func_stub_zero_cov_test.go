//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestSnakeCase_ZeroCov — Pascal/camel → snake (약어 런 처리)
package funcstub

import (
	"strings"
	"testing"
)

func TestRenderFuncStub_ZeroCov(t *testing.T) {
	out := RenderFuncStub("billing", []string{"ChargeCard", "Refund"})
	for _, want := range []string{
		"async def charge_card(",
		"async def refund(",
		"billing.charge_card not implemented",
		"NotImplementedError",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderFuncStub missing %q\n%s", want, out)
		}
	}
}
