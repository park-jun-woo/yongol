//ff:func feature=stml-gen type=test control=sequence
//ff:what zodBaseTypeRecord — object(맵) 값 타입이 z.record(z.<value>()) / 자유형 z.record(z.unknown()) 로 변환되는지 검증
package stml

import "testing"

func TestZodBaseTypeRecord(t *testing.T) {
	if got := zodBaseTypeRecord("string"); got != "z.record(z.string())" {
		t.Errorf("record(string) = %q", got)
	}
	if got := zodBaseTypeRecord("integer"); got != "z.record(z.number().int())" {
		t.Errorf("record(integer) = %q", got)
	}
	if got := zodBaseTypeRecord("*"); got != "z.record(z.unknown())" {
		t.Errorf("record(free-form *) = %q", got)
	}
	if got := zodBaseTypeRecord(""); got != "z.record(z.unknown())" {
		t.Errorf("record(empty) = %q", got)
	}
}
