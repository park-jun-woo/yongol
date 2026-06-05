//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what peek — 현재 토큰을 소비하지 않고 반환 검증

package stml

import "testing"

func TestGuardParserPeek(t *testing.T) {
	toks := []guardToken{
		{kind: tokIdent, text: "a"},
		{kind: tokEOF},
	}

	tests := []struct {
		name     string
		pos      int
		wantKind guardTokKind
		wantText string
	}{
		{name: "peek at first", pos: 0, wantKind: tokIdent, wantText: "a"},
		{name: "peek at eof", pos: 1, wantKind: tokEOF, wantText: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &guardParser{toks: toks, pos: tt.pos}
			got := p.peek()
			if got.kind != tt.wantKind || got.text != tt.wantText {
				t.Errorf("peek() = {%d %q}, want {%d %q}", got.kind, got.text, tt.wantKind, tt.wantText)
			}
			if p.pos != tt.pos {
				t.Errorf("peek mutated pos to %d, want %d", p.pos, tt.pos)
			}
		})
	}
}
