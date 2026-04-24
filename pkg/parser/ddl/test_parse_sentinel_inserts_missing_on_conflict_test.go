//ff:func feature=manifest type=test control=sequence
//ff:what TestParseSentinelInserts_MissingOnConflict — ON CONFLICT 절 없는 INSERT 는 탐지기가 false

package ddl

import (
	"testing"
)

// TestParseSentinelInserts_MissingOnConflict ensures an INSERT without
// the ON CONFLICT DO NOTHING clause is still annotated but the detector
// correctly returns false.
func TestParseSentinelInserts_MissingOnConflict(t *testing.T) {
	content := `-- @sentinel
INSERT INTO t (id, name) VALUES (0, 'sys');
`
	results := parseSentinelInserts(content)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if !results[0].Annotated {
		t.Errorf("annotation expected")
	}
	if sentinelHasOnConflictDoNothing(results[0].SQL) {
		t.Errorf("should NOT detect ON CONFLICT DO NOTHING")
	}
}
