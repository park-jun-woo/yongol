//ff:func feature=cli type=test control=iteration dimension=1
//ff:what TestSplitAdvice — splitAdvice 마커별 분리 검증

package main

import "testing"

func TestSplitAdvice(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantMsg    string
		wantAdvice string
	}{
		{"NoAdvice", "[S-01] bad thing", "[S-01] bad thing", ""},
		{"ArrowAdvice", "[S-01] bad → Advice: fix it", "[S-01] bad", "fix it"},
		{"DashAdvice", "[S-01] bad — Advice: fix it", "[S-01] bad", "fix it"},
		{"NewlineAdvice", "[S-01] bad\n↳ Advice: fix it", "[S-01] bad", "fix it"},
		{"FirstMatchWins", "[S-01] bad\n↳ Advice: first → Advice: second", "[S-01] bad", "first → Advice: second"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, gotAdvice := splitAdvice(tt.input)
			if gotMsg != tt.wantMsg {
				t.Errorf("msg = %q, want %q", gotMsg, tt.wantMsg)
			}
			if gotAdvice != tt.wantAdvice {
				t.Errorf("advice = %q, want %q", gotAdvice, tt.wantAdvice)
			}
		})
	}
}
