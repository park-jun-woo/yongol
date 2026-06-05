//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what zodBaseType — 허용 scalar 집합(string/""/integer/number/boolean)을 panic 없이 변환하는지 검증
package stml

import "testing"

// zodBaseType must accept the allowed scalar set (incl. string/"") without panic.
func TestZodBaseType_AllowedScalars(t *testing.T) {
	cases := map[string]string{
		"string":  "z.string()",
		"":        "z.string()",
		"integer": "z.number().int()",
		"number":  "z.number()",
		"boolean": "z.boolean()",
	}
	for typ, want := range cases {
		if got := zodBaseType(typ); got != want {
			t.Errorf("zodBaseType(%q) = %q, want %q", typ, got, want)
		}
	}
}
