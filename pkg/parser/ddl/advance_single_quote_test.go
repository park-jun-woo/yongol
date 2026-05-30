//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what advanceSingleQuote — '' 이스케이프 / 인용 토글 / 라인 끝 경계 검증

package ddl

import "testing"

func TestAdvanceSingleQuote(t *testing.T) {
	cases := []struct {
		name    string
		line    string
		i       int
		inSQ    bool
		wantSQ  bool
		wantPos int
	}{
		{"open quote", "'abc'", 0, false, true, 1},
		{"close quote", "'abc'", 4, true, false, 5},
		{"escape pair keeps inSQ", "''", 0, true, true, 2},
		{"quote at end not escape", "a'", 1, true, false, 2},
		{"lone quote toggles on", "x'y", 1, false, true, 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotSQ, gotPos := advanceSingleQuote(c.line, c.i, c.inSQ)
			if gotSQ != c.wantSQ || gotPos != c.wantPos {
				t.Errorf("advanceSingleQuote(%q,%d,%v) = (%v,%d), want (%v,%d)",
					c.line, c.i, c.inSQ, gotSQ, gotPos, c.wantSQ, c.wantPos)
			}
		})
	}
}
