//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what splitTrailingComment — `-- ` 주석 분리 / 문자열 리터럴 내 `--` 무시

package ddl

import "testing"

func TestSplitTrailingComment(t *testing.T) {
	cases := []struct {
		name        string
		line        string
		wantDDL     string
		wantComment string
	}{
		{"plain comment", "id BIGINT -- @nullable", "id BIGINT ", "@nullable"},
		{"no comment", "id BIGINT NOT NULL", "id BIGINT NOT NULL", ""},
		{"dash inside string literal", "DEFAULT 'a--b'", "DEFAULT 'a--b'", ""},
		{"comment after string", "DEFAULT 'x' -- note", "DEFAULT 'x' ", "note"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ddl, comment := splitTrailingComment(c.line)
			if ddl != c.wantDDL || comment != c.wantComment {
				t.Errorf("splitTrailingComment(%q) = (%q,%q), want (%q,%q)",
					c.line, ddl, comment, c.wantDDL, c.wantComment)
			}
		})
	}
}
