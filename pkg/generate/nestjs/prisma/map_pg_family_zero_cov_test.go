//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"
)

func TestMapPGFamily_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"BIGINT": "Int", "INTEGER": "Int", "SMALLINT": "Int",
		"TEXT": "String", "CITEXT": "String",
		"BOOLEAN": "Boolean", "BOOL": "Boolean",
		"TIMESTAMP": "DateTime", "DATE": "DateTime",
		"UUID":  "String",
		"JSONB": "Json", "JSON": "Json",
		"NUMERIC": "Decimal", "DECIMAL": "Decimal",
		"FLOAT": "Float", "REAL": "Float",
		"BYTEA": "Bytes",
		"INET":  "String", "INTERVAL": "String",
		"UNKNOWNTYPE": "String",
	}
	for in, want := range cases {
		if got := mapPGFamily(in); got != want {
			t.Errorf("mapPGFamily(%q)=%q want %q", in, got, want)
		}
	}
}
