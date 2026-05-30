//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ScanSentinelInserts + SentinelHasOnConflictDoNothing — 공개 sentinel API

package ddl

import "testing"

func TestScanSentinelInserts(t *testing.T) {
	content := `-- @sentinel
INSERT INTO roles (id, name) VALUES (1, 'admin') ON CONFLICT DO NOTHING;
INSERT INTO other (id) VALUES (2);`
	got := ScanSentinelInserts(content)
	if len(got) == 0 {
		t.Fatal("expected at least one sentinel scan result")
	}
	// The annotated INSERT must be flagged.
	var found bool
	for _, r := range got {
		if r.Table == "roles" {
			found = true
			if !r.Annotated {
				t.Errorf("roles INSERT should be Annotated")
			}
			if r.StartLine <= 0 {
				t.Errorf("StartLine = %d, want > 0", r.StartLine)
			}
		}
	}
	if !found {
		t.Errorf("roles sentinel not found in %+v", got)
	}
}

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
