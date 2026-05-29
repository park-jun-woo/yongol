//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what slogAttrLine — 모든 Go 타입을 slog.Any 로 통일 emit (BUG-024)

package sqlcpost

import "testing"

// TestSlogAttrLine_AlwaysAny pins the post-BUG-024 behaviour: every
// non-redacted column goes through slog.Any, regardless of declared Go
// type. The previous type-dispatch (slog.String / slog.Time / slog.Int64 /
// …) failed to compile when sqlc's pgx/v5 sql_package introduced pgtype
// wrappers. We intentionally lost that type variety to gain pgtype
// compatibility — do not re-introduce typed constructors without
// addressing BUG-024.
func TestSlogAttrLine_AlwaysAny(t *testing.T) {
	cases := []struct {
		name   string
		goType string
	}{
		{"int64 column", "int64"},
		{"string column", "string"},
		{"bool column", "bool"},
		{"time.Time column", "time.Time"},
		{"float64 column", "float64"},
		{"json.RawMessage column", "json.RawMessage"},
		{"pgtype.UUID column", "pgtype.UUID"},
		{"pgtype.Timestamp column", "pgtype.Timestamp"},
		{"pgtype.Timestamptz column", "pgtype.Timestamptz"},
		{"unknown type column", "somepkg.CustomType"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := slogAttrLine("my_col", "MyCol", c.goType)
			want := `slog.Any("my_col", r.MyCol)`
			if got != want {
				t.Fatalf("goType=%s:\n  got:  %s\n  want: %s", c.goType, got, want)
			}
		})
	}
}
