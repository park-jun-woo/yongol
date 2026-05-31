//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"
)

func TestPyFamily(t *testing.T) {
	cases := map[string]string{
		"BIGINT":           "int",
		"INT2":             "int",
		"TEXT":             "str",
		"CITEXT":           "str",
		"BOOLEAN":          "bool",
		"TIMESTAMPTZ":      "datetime",
		"DATE":             "date",
		"UUID":             "uuid.UUID",
		"JSONB":            "dict[str, Any]",
		"NUMERIC":          "Decimal",
		"DOUBLE PRECISION": "float",
		"BYTEA":            "bytes",
		"INET":             "str",
		"INTERVAL":         "timedelta",
		"UNKNOWNTYPE":      "str",
	}
	for in, want := range cases {
		if got := pyFamily(in); got != want {
			t.Errorf("pyFamily(%q) = %q, want %q", in, got, want)
		}
	}
}
