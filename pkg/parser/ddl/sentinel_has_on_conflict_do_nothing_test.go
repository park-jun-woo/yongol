//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ScanSentinelInserts + SentinelHasOnConflictDoNothing — 공개 sentinel API
package ddl

import (
	"testing"
)

func TestSentinelHasOnConflictDoNothing(t *testing.T) {
	cases := []struct {
		sql  string
		want bool
	}{
		{"INSERT INTO t VALUES (1) ON CONFLICT DO NOTHING;", true},
		{"INSERT INTO t (id) VALUES (1) ON CONFLICT (id) DO NOTHING;", true},
		{"INSERT INTO t VALUES (1);", false},
		{"INSERT INTO t VALUES (1) ON CONFLICT (id) DO UPDATE SET x=1;", false},
	}
	for _, c := range cases {
		if got := SentinelHasOnConflictDoNothing(c.sql); got != c.want {
			t.Errorf("SentinelHasOnConflictDoNothing(%q) = %v, want %v", c.sql, got, c.want)
		}
	}
}
