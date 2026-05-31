//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestShouldParseSQL — 디렉토리/비-SQL/skip/baseline 제외 규칙
package migration

import (
	"testing"
)

func TestShouldParseSQL(t *testing.T) {
	skip := map[string]bool{"skipme.sql": true}
	cases := []struct {
		name  string
		isDir bool
		fname string
		want  bool
	}{
		{"plain sql", false, "schema.sql", true},
		{"uppercase ext", false, "SCHEMA.SQL", true},
		{"directory", true, "subdir.sql", false},
		{"non-sql", false, "notes.txt", false},
		{"skipped", false, "skipme.sql", false},
		{"snapshot baseline", false, SnapshotFileName, false},
		{"legacy baseline", false, LegacySnapshotFileName, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := shouldParseSQL(c.isDir, c.fname, skip); got != c.want {
				t.Errorf("shouldParseSQL(%v,%q) = %v, want %v", c.isDir, c.fname, got, c.want)
			}
		})
	}
}
