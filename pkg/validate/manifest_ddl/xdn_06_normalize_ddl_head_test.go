//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what normalizeDDLHead — 배열/파라미터 제거 + 대문자 정규화 검증

package manifest_ddl

import "testing"

func TestNormalizeDDLHead(t *testing.T) {
	tests := []struct {
		name    string
		rawType string
		want    string
	}{
		{
			name:    "simple type uppercase",
			rawType: "BIGINT",
			want:    "BIGINT",
		},
		{
			name:    "lowercase normalized to upper",
			rawType: "bigint",
			want:    "BIGINT",
		},
		{
			name:    "varchar with length parameter stripped",
			rawType: "VARCHAR(255)",
			want:    "VARCHAR",
		},
		{
			name:    "array suffix removed",
			rawType: "TEXT[]",
			want:    "TEXT",
		},
		{
			name:    "array with parameter strips both",
			rawType: "VARCHAR(100)[]",
			want:    "VARCHAR",
		},
		{
			name:    "leading/trailing spaces trimmed",
			rawType: "  INTEGER  ",
			want:    "INTEGER",
		},
		{
			name:    "character varying",
			rawType: "CHARACTER VARYING",
			want:    "CHARACTER VARYING",
		},
		{
			name:    "numeric with precision",
			rawType: "NUMERIC(10,2)",
			want:    "NUMERIC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeDDLHead(tt.rawType)
			if got != tt.want {
				t.Errorf("normalizeDDLHead(%q) = %q, want %q", tt.rawType, got, tt.want)
			}
		})
	}
}
