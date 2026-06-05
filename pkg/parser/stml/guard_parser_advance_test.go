//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what advance — 현재 토큰 소비·반환, EOF에서 위치 고정 검증

package stml

import "testing"

func TestGuardParserAdvance(t *testing.T) {
	toks := []guardToken{
		{kind: tokIdent, text: "a"},
		{kind: tokDot, text: "."},
		{kind: tokEOF},
	}

	tests := []struct {
		name     string
		wantText string
		wantPos  int
	}{
		{name: "first advance returns a", wantText: "a", wantPos: 1},
		{name: "second advance returns dot", wantText: ".", wantPos: 2},
		{name: "third advance returns EOF, pos stays", wantText: "", wantPos: 2},
		{name: "fourth advance at EOF, pos stays", wantText: "", wantPos: 2},
	}

	p := &guardParser{toks: toks}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p.advance()
			if got.text != tt.wantText {
				t.Errorf("advance() text = %q, want %q", got.text, tt.wantText)
			}
			if p.pos != tt.wantPos {
				t.Errorf("pos = %d, want %d", p.pos, tt.wantPos)
			}
		})
	}
}
