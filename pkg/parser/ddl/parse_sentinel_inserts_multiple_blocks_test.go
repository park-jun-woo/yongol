//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what TestParseSentinelInserts_MultipleBlocks — 다수 @sentinel 블록이 독립적으로 수집

package ddl

import (
	"testing"
)

// TestParseSentinelInserts_MultipleBlocks asserts that multiple
// `-- @sentinel` INSERT blocks are each captured with Annotated=true.
func TestParseSentinelInserts_MultipleBlocks(t *testing.T) {
	content := `CREATE TABLE lookups (id BIGINT, code VARCHAR(8));

-- @sentinel
INSERT INTO lookups (id, code) VALUES (0, 'unknown') ON CONFLICT DO NOTHING;

-- @sentinel
INSERT INTO lookups (id, code) VALUES (1, 'active') ON CONFLICT DO NOTHING;
`
	results := parseSentinelInserts(content)
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	for i, r := range results {
		if !r.Annotated {
			t.Errorf("block %d: expected Annotated=true", i)
		}
	}
}
