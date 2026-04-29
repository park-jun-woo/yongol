//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what extractRawType — 다중 단어 PG 타입 보존 + 단일 토큰 fallback 회귀

package ddl

import "testing"

func TestExtractRawType(t *testing.T) {
	cases := []struct {
		name     string
		tokens   []string
		wantRaw  string
		wantUsed int
	}{
		{
			name:     "single_token_with_constraints",
			tokens:   []string{"BIGINT", "NOT", "NULL"},
			wantRaw:  "BIGINT",
			wantUsed: 1,
		},
		{
			name:     "single_token_with_param",
			tokens:   []string{"VARCHAR(255)"},
			wantRaw:  "VARCHAR(255)",
			wantUsed: 1,
		},
		{
			name:     "double_precision",
			tokens:   []string{"DOUBLE", "PRECISION", "NOT", "NULL"},
			wantRaw:  "DOUBLE PRECISION",
			wantUsed: 2,
		},
		{
			name:     "timestamp_with_time_zone",
			tokens:   []string{"TIMESTAMP", "WITH", "TIME", "ZONE"},
			wantRaw:  "TIMESTAMP WITH TIME ZONE",
			wantUsed: 4,
		},
		{
			name:     "timestamp_without_time_zone",
			tokens:   []string{"TIMESTAMP", "WITHOUT", "TIME", "ZONE"},
			wantRaw:  "TIMESTAMP WITHOUT TIME ZONE",
			wantUsed: 4,
		},
		{
			name:     "character_varying_with_param",
			tokens:   []string{"CHARACTER", "VARYING(255)"},
			wantRaw:  "CHARACTER VARYING(255)",
			wantUsed: 2,
		},
		{
			name:     "character_with_param",
			tokens:   []string{"CHARACTER(10)"},
			wantRaw:  "CHARACTER(10)",
			wantUsed: 1,
		},
		{
			name:     "bit_varying_with_param",
			tokens:   []string{"BIT", "VARYING(8)"},
			wantRaw:  "BIT VARYING(8)",
			wantUsed: 2,
		},
		{
			name:     "lonely_timestamp",
			tokens:   []string{"TIMESTAMP"},
			wantRaw:  "TIMESTAMP",
			wantUsed: 1,
		},
		{
			name:     "empty_tokens",
			tokens:   []string{},
			wantRaw:  "",
			wantUsed: 0,
		},
		{
			name:     "time_with_time_zone",
			tokens:   []string{"TIME", "WITH", "TIME", "ZONE", "NOT", "NULL"},
			wantRaw:  "TIME WITH TIME ZONE",
			wantUsed: 4,
		},
		{
			name:     "lowercase_double_precision",
			tokens:   []string{"double", "precision"},
			wantRaw:  "DOUBLE PRECISION",
			wantUsed: 2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotRaw, gotUsed := extractRawType(c.tokens)
			if gotRaw != c.wantRaw {
				t.Errorf("raw = %q, want %q", gotRaw, c.wantRaw)
			}
			if gotUsed != c.wantUsed {
				t.Errorf("used = %d, want %d", gotUsed, c.wantUsed)
			}
		})
	}
}
