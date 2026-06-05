//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what lexGuardString — 따옴표 문자열 리터럴 토큰 읽기 (정상/미종료) 검증

package stml

import "testing"

func TestLexGuardString(t *testing.T) {
	tests := []struct {
		name     string
		runes    string
		i        int
		wantText string
		wantNext int
		wantErr  bool
	}{
		{name: "single quoted", runes: "'active'", i: 0, wantText: "active", wantNext: 8},
		{name: "double quoted", runes: `"draft"`, i: 0, wantText: "draft", wantNext: 7},
		{name: "empty string", runes: "''", i: 0, wantText: "", wantNext: 2},
		{name: "quote mid-string", runes: "x='hi'", i: 2, wantText: "hi", wantNext: 6},
		{name: "unterminated", runes: "'open", i: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertLexGuardString(t, tt.runes, tt.i, tt.wantText, tt.wantNext, tt.wantErr)
		})
	}
}
