//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what shouldKeepImport — 단일 import 라인 유지 여부 결정 (blank / 사용 여부 분기)

package boot

import "testing"

func TestShouldKeepImport(t *testing.T) {
	cases := []struct {
		name      string
		imp       string
		body      string
		keepBlank bool
		want      bool
	}{
		{"used package", `"strconv"`, "n := strconv.Atoi(x)", false, true},
		{"unused package", `"strconv"`, "n := other.Atoi(x)", false, false},
		{"blank import kept", `_ "github.com/lib/pq"`, "no usage here", true, true},
		{"blank import dropped", `_ "github.com/lib/pq"`, "no usage here", false, false},
		{"word boundary blocks suffix match", `"database/sql"`, "x := otelsql.Open()", false, false},
		{"word boundary allows real use", `"database/sql"`, "var d sql.DB", false, true},
	}
	for _, c := range cases {
		if got := shouldKeepImport(c.imp, c.body, c.keepBlank); got != c.want {
			t.Errorf("%s: shouldKeepImport(%q,...) = %v, want %v", c.name, c.imp, got, c.want)
		}
	}
}
