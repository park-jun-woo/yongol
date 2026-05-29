//ff:func feature=manifest type=test control=sequence
//ff:what TestParseSentinelInserts_UnannotatedStillScanned — @sentinel 없어도 INSERT 는 여전히 수집되지만 Annotated=false

package ddl

import (
	"testing"
)

// TestParseSentinelInserts_UnannotatedStillScanned verifies that a plain
// INSERT (no `-- @sentinel` above it) is still returned by the scanner,
// but flagged Annotated=false so the migration emitter skips it.
func TestParseSentinelInserts_UnannotatedStillScanned(t *testing.T) {
	content := `CREATE TABLE t (id BIGINT);

INSERT INTO t (id) VALUES (0) ON CONFLICT DO NOTHING;
`
	results := parseSentinelInserts(content)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Annotated {
		t.Errorf("expected Annotated=false when no -- @sentinel above INSERT")
	}
}
