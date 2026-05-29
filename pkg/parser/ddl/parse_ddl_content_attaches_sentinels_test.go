//ff:func feature=manifest type=test control=sequence
//ff:what TestParseDDLContent_AttachesSentinels — parseDDLContent 가 Table.Sentinels 에 INSERT 블록 첨부

package ddl

import (
	"strings"
	"testing"
)

// TestParseDDLContent_AttachesSentinels exercises the DDL content parser
// to ensure `-- @sentinel` INSERT blocks end up on the target table's
// Sentinels slice with the correct source File attribution.
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
