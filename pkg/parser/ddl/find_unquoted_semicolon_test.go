//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what findUnquotedSemicolon — 인용 밖 첫 `;` 위치 / 리터럴 내 `;` 무시 / '' 이스케이프

package ddl

import "testing"

func TestFindUnquotedSemicolon(t *testing.T) {
	cases := []struct {
		name     string
		line     string
		inSingle bool
		wantPos  int
		wantOK   bool
	}{
		{"plain semicolon", "VALUES (1);", false, 10, true},
		{"semicolon inside literal", "VALUES ('a;b'", false, -1, false},
		{"semicolon after literal closes", "'a;b');", false, 6, true},
		{"no semicolon", "id BIGINT", false, -1, false},
		{"start inside literal", "abc' ;", true, 5, true},
		{"escaped quote then semicolon", "''';", true, 3, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pos, ok := findUnquotedSemicolon(c.line, c.inSingle)
			if pos != c.wantPos || ok != c.wantOK {
				t.Errorf("findUnquotedSemicolon(%q,%v) = (%d,%v), want (%d,%v)",
					c.line, c.inSingle, pos, ok, c.wantPos, c.wantOK)
			}
		})
	}
}
