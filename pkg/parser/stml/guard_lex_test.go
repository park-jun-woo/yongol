//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what lexGuard — 가드 조건 문자열을 어휘 토큰 슬라이스로 분해 (정상/에러) 검증

package stml

import "testing"

func TestLexGuard(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []guardToken
		wantErr bool
	}{
		{
			name:  "compare and combination",
			input: "a.b=active && c.d!=draft",
			want: []guardToken{
				{kind: tokIdent, text: "a"},
				{kind: tokDot, text: "."},
				{kind: tokIdent, text: "b"},
				{kind: tokOp, text: "="},
				{kind: tokIdent, text: "active"},
				{kind: tokAnd, text: "&&"},
				{kind: tokIdent, text: "c"},
				{kind: tokDot, text: "."},
				{kind: tokIdent, text: "d"},
				{kind: tokOp, text: "!="},
				{kind: tokIdent, text: "draft"},
				{kind: tokEOF},
			},
		},
		{
			name:  "negation group and quoted string",
			input: "!(x.y='hi')",
			want: []guardToken{
				{kind: tokNot, text: "!"},
				{kind: tokLParen, text: "("},
				{kind: tokIdent, text: "x"},
				{kind: tokDot, text: "."},
				{kind: tokIdent, text: "y"},
				{kind: tokOp, text: "="},
				{kind: tokString, text: "hi"},
				{kind: tokRParen, text: ")"},
				{kind: tokEOF},
			},
		},
		{
			name:  "empty produces only EOF",
			input: "   ",
			want: []guardToken{
				{kind: tokEOF},
			},
		},
		{
			name:    "illegal character",
			input:   "a.b @ c",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertLexGuard(t, tt.input, tt.want, tt.wantErr)
		})
	}
}
