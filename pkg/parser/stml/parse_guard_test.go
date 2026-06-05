//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what ParseGuard — 가드 조건 문자열 EBNF 파싱 (정상/렉스 에러/파스 에러/잔여 토큰 에러) 검증

package stml

import "testing"

func TestParseGuard(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantKind GuardKind
	}{
		{name: "single compare", input: "a.b=x", wantKind: GuardCompare},
		{name: "binary combination", input: "a.b=x && c.d=y", wantKind: GuardBinary},
		{name: "lex error", input: "a.b=x @ y", wantErr: true},
		{name: "parse error", input: "&&", wantErr: true},
		{name: "trailing token", input: "a.b=x c.d=y", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuard(t, tt.input, tt.wantErr, tt.wantKind)
		})
	}
}
