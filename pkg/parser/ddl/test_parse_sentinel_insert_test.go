//ff:func feature=manifest type=test control=sequence
//ff:what parseSentinelInserts — @sentinel annotated INSERT collection and ON CONFLICT DO NOTHING detection

package ddl

import (
	"strings"
	"testing"
)

func TestIsSentinelAnnotation(t *testing.T) {
	if !isSentinelAnnotation("-- @sentinel") {
		t.Errorf("missed -- @sentinel")
	}
	if isSentinelAnnotation("-- sentinel") {
		t.Errorf("false positive without @")
	}
	if isSentinelAnnotation("INSERT INTO t VALUES (0);") {
		t.Errorf("false positive for non-comment")
	}
}

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

func TestParseSentinelInserts_SemicolonInLiteral(t *testing.T) {
	content := `-- @sentinel
INSERT INTO t (id, msg) VALUES (0, 'hi; there') ON CONFLICT DO NOTHING;
`
	results := parseSentinelInserts(content)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if !strings.Contains(results[0].SQL, "hi; there") {
		t.Errorf("literal semicolon was lost: %q", results[0].SQL)
	}
	if !strings.HasSuffix(strings.TrimSpace(results[0].SQL), ";") {
		t.Errorf("terminator ; missing: %q", results[0].SQL)
	}
}

func TestParseDDLContent_AttachesSentinels(t *testing.T) {
	content := `CREATE TABLE organizations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- @sentinel
INSERT INTO organizations (id, name) VALUES (0, 'sys') ON CONFLICT DO NOTHING;
`
	tables := map[string]*Table{}
	parseDDLContent(content, tables, "organizations.sql")
	tbl := tables["organizations"]
	if tbl == nil {
		t.Fatalf("organizations table not parsed")
	}
	if len(tbl.Sentinels) != 1 {
		t.Fatalf("expected 1 sentinel attached, got %d", len(tbl.Sentinels))
	}
	if !strings.Contains(tbl.Sentinels[0].SQL, "INSERT INTO organizations") {
		t.Errorf("sentinel SQL mismatch: %q", tbl.Sentinels[0].SQL)
	}
	if tbl.Sentinels[0].File != "organizations.sql" {
		t.Errorf("sentinel File field not set correctly: %q", tbl.Sentinels[0].File)
	}
}
