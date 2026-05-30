package funcstub

import (
	"strings"
	"testing"
)

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderFuncModule_ZeroCov — 외부 패키지 stub module 소스 생성

func TestRenderFuncModule_ZeroCov(t *testing.T) {
	out := RenderFuncModule("billing", []string{"Charge"})
	for _, want := range []string{"BillingService", "export class BillingModule", "./billing.service", "@Global()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderFuncModule missing %q\n%s", want, out)
		}
	}
}

//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderFuncService_ZeroCov — 외부 패키지 stub service 메서드 생성

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
