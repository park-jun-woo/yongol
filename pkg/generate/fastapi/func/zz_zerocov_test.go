package funcstub

import (
	"strings"
	"testing"
)

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestSnakeCase_ZeroCov — Pascal/camel → snake (약어 런 처리)

func TestSnakeCase_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"":            "",
		"Charge":      "charge",
		"ChargeCard":  "charge_card",
		"getURL":      "get_url",
		"URLParser":   "url_parser",
		"parseID":     "parse_id",
	}
	for in, want := range cases {
		if got := snakeCase(in); got != want {
			t.Errorf("snakeCase(%q)=%q want %q", in, got, want)
		}
	}
}

//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderFuncStub_ZeroCov — Python stub 모듈 메서드 생성

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
