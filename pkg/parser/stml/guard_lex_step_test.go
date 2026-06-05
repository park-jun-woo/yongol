//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what lexGuardStep — 인덱스 i에서 토큰 하나 분류·추출 (공백/문자열/심볼/식별자/에러) 검증

package stml

import "testing"

func TestLexGuardStep(t *testing.T) {
	tests := []struct {
		name     string
		runes    string
		i        int
		wantKind guardTokKind
		wantText string
		wantEmit bool
		wantNext int
		wantErr  bool
	}{
		{name: "whitespace skipped", runes: " a", i: 0, wantEmit: false, wantNext: 1},
		{name: "single quoted string", runes: "'hi'", i: 0, wantKind: tokString, wantText: "hi", wantEmit: true, wantNext: 4},
		{name: "symbol and", runes: "&&x", i: 0, wantKind: tokAnd, wantText: "&&", wantEmit: true, wantNext: 2},
		{name: "identifier", runes: "status=", i: 0, wantKind: tokIdent, wantText: "status", wantEmit: true, wantNext: 6},
		{name: "illegal char", runes: "@", i: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertLexGuardStep(t, tt.runes, tt.i, tt.wantKind, tt.wantText, tt.wantEmit, tt.wantNext, tt.wantErr)
		})
	}
}
