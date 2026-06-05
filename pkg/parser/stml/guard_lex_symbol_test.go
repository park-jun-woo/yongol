//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what lexGuardSymbol — 연산자/논리/그룹 심볼 토큰 읽기 (&& || ! != ( ) . = > < >= <=, 에러) 검증

package stml

import "testing"

func TestLexGuardSymbol(t *testing.T) {
	tests := []struct {
		name     string
		runes    string
		i        int
		wantKind guardTokKind
		wantText string
		wantNext int
		wantErr  bool
	}{
		{name: "and", runes: "&&x", i: 0, wantKind: tokAnd, wantText: "&&", wantNext: 2},
		{name: "lone ampersand error", runes: "&x", i: 0, wantErr: true},
		{name: "or", runes: "||x", i: 0, wantKind: tokOr, wantText: "||", wantNext: 2},
		{name: "lone pipe error", runes: "|x", i: 0, wantErr: true},
		{name: "not equal", runes: "!=v", i: 0, wantKind: tokOp, wantText: "!=", wantNext: 2},
		{name: "not", runes: "!x", i: 0, wantKind: tokNot, wantText: "!", wantNext: 1},
		{name: "lparen", runes: "(", i: 0, wantKind: tokLParen, wantText: "(", wantNext: 1},
		{name: "rparen", runes: ")", i: 0, wantKind: tokRParen, wantText: ")", wantNext: 1},
		{name: "dot", runes: ".", i: 0, wantKind: tokDot, wantText: ".", wantNext: 1},
		{name: "equals", runes: "=v", i: 0, wantKind: tokOp, wantText: "=", wantNext: 1},
		{name: "greater equal", runes: ">=3", i: 0, wantKind: tokOp, wantText: ">=", wantNext: 2},
		{name: "greater", runes: ">3", i: 0, wantKind: tokOp, wantText: ">", wantNext: 1},
		{name: "less equal", runes: "<=3", i: 0, wantKind: tokOp, wantText: "<=", wantNext: 2},
		{name: "less", runes: "<3", i: 0, wantKind: tokOp, wantText: "<", wantNext: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertLexGuardSymbol(t, tt.runes, tt.i, tt.wantKind, tt.wantText, tt.wantNext, tt.wantErr)
		})
	}
}
