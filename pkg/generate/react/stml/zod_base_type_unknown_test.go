//ff:func feature=stml-gen type=test control=sequence
//ff:what zodBaseType — 미지원 scalar 타입이 silent z.string() 폴백 대신 *zodGenError panic 으로 드러나는지 검증
package stml

import "testing"

// An unknown scalar type panics with *zodGenError instead of degrading to string.
func TestZodBaseType_UnknownPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on unsupported type, got none")
		}
		ze, ok := r.(*zodGenError)
		if !ok {
			t.Fatalf("expected *zodGenError, got %T", r)
		}
		if ze.Type != "weirdtype" {
			t.Errorf("zodGenError.Type = %q, want \"weirdtype\"", ze.Type)
		}
	}()
	_ = zodBaseType("weirdtype")
}
