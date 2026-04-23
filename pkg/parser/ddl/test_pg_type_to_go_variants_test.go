//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what pgTypeToGo — 주요 PostgreSQL 타입 → Go 타입 매핑 회귀

package ddl

import "testing"

func TestPgTypeToGo_Variants(t *testing.T) {
	cases := map[string]string{
		"BIGINT":      "int64",
		"SERIAL":      "int64",
		"TEXT":        "string",
		"BOOLEAN":     "bool",
		"TIMESTAMPTZ": "time.Time",
		"JSONB":       "json.RawMessage",
		"NUMERIC":     "float64",
		"VARCHAR(64)": "string",
		"UNKNOWN_XY":  "string",
	}
	for in, want := range cases {
		if got := pgTypeToGo(in); got != want {
			t.Errorf("pgTypeToGo(%q) = %q, want %q", in, got, want)
		}
	}
}
