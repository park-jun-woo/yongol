//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what lexGuardIdent — 식별자/숫자/enum-literal 토큰 읽기 검증 (소비 길이 포함)

package stml

import "testing"

func TestLexGuardIdent(t *testing.T) {
	tests := []struct {
		name     string
		runes    string
		start    int
		wantText string
		wantNext int
	}{
		{name: "plain identifier", runes: "status", start: 0, wantText: "status", wantNext: 6},
		{name: "stops at dot", runes: "model.field", start: 0, wantText: "model", wantNext: 5},
		{name: "number", runes: "123 ", start: 0, wantText: "123", wantNext: 3},
		{name: "enum literal with hyphen and underscore", runes: "in_progress-2=x", start: 0, wantText: "in_progress-2", wantNext: 13},
		{name: "mid-string start", runes: "a.bcd", start: 2, wantText: "bcd", wantNext: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertLexGuardIdent(t, tt.runes, tt.start, tt.wantText, tt.wantNext)
		})
	}
}
