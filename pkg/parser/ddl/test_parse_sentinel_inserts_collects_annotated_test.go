//ff:func feature=manifest type=test control=sequence
//ff:what TestParseSentinelInserts_CollectsAnnotated — @sentinel annotated INSERT 1개 수집 + ON CONFLICT DO NOTHING 인식

package ddl

import (
	"strings"
	"testing"
)

// TestParseSentinelInserts_CollectsAnnotated asserts the happy path: a
// single `-- @sentinel` INSERT block is captured with Annotated=true, a
// verbatim SQL body and ON CONFLICT DO NOTHING detection.
func TestParseSentinelInserts_CollectsAnnotated(t *testing.T) {
	content := `CREATE TABLE organizations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- @sentinel
INSERT INTO organizations (id, name)
OVERRIDING SYSTEM VALUE
VALUES (0, 'system')
ON CONFLICT DO NOTHING;
`
	results := parseSentinelInserts(content)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	r := results[0]
	if r.Table != "organizations" {
		t.Errorf("table: got %q, want organizations", r.Table)
	}
	if !r.Annotated {
		t.Errorf("expected Annotated=true")
	}
	if !strings.HasPrefix(strings.TrimSpace(r.SQL), "INSERT INTO organizations") {
		t.Errorf("SQL prefix mismatch: %q", r.SQL)
	}
	if !strings.HasSuffix(strings.TrimSpace(r.SQL), ";") {
		t.Errorf("SQL not terminated by ;: %q", r.SQL)
	}
	if !sentinelHasOnConflictDoNothing(r.SQL) {
		t.Errorf("ON CONFLICT DO NOTHING should be detected")
	}
}
