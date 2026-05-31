//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"
)

func TestMapSAFamily(t *testing.T) {
	cases := map[string]string{
		"BIGINT":                      "Integer",
		"INTEGER":                     "Integer",
		"TEXT":                        "Text",
		"VARCHAR":                     "String",
		"BOOLEAN":                     "Boolean",
		"TIMESTAMP":                   "DateTime(timezone=False)",
		"TIMESTAMP WITHOUT TIME ZONE": "DateTime(timezone=False)",
		"TIMESTAMPTZ":                 "DateTime(timezone=True)",
		"DATE":                        "Date",
		"UUID":                        "Uuid",
		"JSONB":                       "JSONB",
		"NUMERIC":                     "Numeric",
		"REAL":                        "Float",
		"BYTEA":                       "LargeBinary",
		"CIDR":                        "INET",
		"INTERVAL":                    "Interval",
		"WEIRD":                       "String",
	}
	for in, want := range cases {
		if got := mapSAFamily(in); got != want {
			t.Errorf("mapSAFamily(%q) = %q, want %q", in, got, want)
		}
	}
}
