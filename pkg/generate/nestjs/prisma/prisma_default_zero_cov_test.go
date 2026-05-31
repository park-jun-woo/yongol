//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestPrismaDefault_ZeroCov(t *testing.T) {
	cases := []struct {
		lit  string
		want string
	}{
		{"", `""`},
		{"now()", "now()"},
		{"CURRENT_TIMESTAMP", "now()"},
		{"gen_random_uuid()", "uuid()"},
		{"uuid_generate_v4()", "uuid()"},
		{"TRUE", "true"},
		{"false", "false"},
		{"42", "42"},
		{"3.14", "3.14"},
		{"hello", `"hello"`},
	}
	for _, c := range cases {
		col := ddl.Column{DefaultLiteral: c.lit}
		if got := prismaDefault(col); got != c.want {
			t.Errorf("prismaDefault(%q)=%q want %q", c.lit, got, c.want)
		}
	}
}
