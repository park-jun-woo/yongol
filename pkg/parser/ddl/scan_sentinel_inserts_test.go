//ff:func feature=manifest type=test control=sequence
//ff:what ScanSentinelInserts + SentinelHasOnConflictDoNothing — 공개 sentinel API
package ddl

import (
	"testing"
)

func TestScanSentinelInserts(t *testing.T) {
	content := `-- @sentinel
INSERT INTO roles (id, name) VALUES (1, 'admin') ON CONFLICT DO NOTHING;
INSERT INTO other (id) VALUES (2);`
	got := ScanSentinelInserts(content)
	if len(got) == 0 {
		t.Fatal("expected at least one sentinel scan result")
	}
	if !rolesSentinelAnnotated(t, got) {
		t.Errorf("roles sentinel not found in %+v", got)
	}
}
