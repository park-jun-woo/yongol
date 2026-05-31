//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"
)

func TestPgToPrismaType_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"BIGINT":        "Int",
		"varchar(255)":  "String",
		"text[]":        "String[]",
		"numeric(10,2)": "Decimal",
	}
	for in, want := range cases {
		if got := pgToPrismaType(in); got != want {
			t.Errorf("pgToPrismaType(%q)=%q want %q", in, got, want)
		}
	}
}
