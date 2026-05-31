//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"testing"
)

func TestFormatCallTarget(t *testing.T) {
	if got := formatCallTarget("", "HoldEscrow"); got != "holdEscrow" {
		t.Errorf("local = %q", got)
	}
	if got := formatCallTarget("billing", "HoldEscrow"); got != "this.billingService.holdEscrow" {
		t.Errorf("di = %q", got)
	}
}
