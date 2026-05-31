//ff:func feature=migration type=test control=iteration dimension=1
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseCreateIndex(t *testing.T) {
	s := NewSchema()
	cases := []struct {
		name    string
		stmt    string
		idxName string
		unique  bool
		method  string
		where   string
		cols    []string
	}{
		{"plain", "CREATE INDEX idx_a ON users (email)", "idx_a", false, "", "", []string{"email"}},
		{"unique", "CREATE UNIQUE INDEX idx_u ON users (email)", "idx_u", true, "", "", []string{"email"}},
		{"using gin", "CREATE INDEX idx_g ON users USING gin (doc)", "idx_g", false, "gin", "", []string{"doc"}},
		{"partial", "CREATE INDEX idx_w ON users (email) WHERE email IS NOT NULL", "idx_w", false, "", "email IS NOT NULL", []string{"email"}},
		{"multi col", "CREATE INDEX idx_m ON users (a, b)", "idx_m", false, "", "", []string{"a", "b"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertParseCreateIndex(t, s, c.stmt, c.idxName, c.unique, c.method, c.where, c.cols)
		})
	}
}
