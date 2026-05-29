//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what TestParseTables_CreateIndexUsing — CREATE INDEX ... USING <method> 절 보존 (BUG-032 / Phase009)

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_CreateIndex_UsingMethod(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE refresh_tokens (
    id BIGSERIAL PRIMARY KEY,
    claims JSONB
);
CREATE INDEX refresh_tokens_claims_idx ON refresh_tokens USING GIN (claims);
CREATE INDEX idx_refresh_tokens_id ON refresh_tokens USING BTREE (id);
CREATE INDEX idx_refresh_tokens_trgm ON refresh_tokens USING HASH (id);
CREATE INDEX idx_refresh_tokens_default ON refresh_tokens (id);
`
	if err := os.WriteFile(filepath.Join(dir, "schema.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatalf("tables count = %d", len(tables))
	}
	tb := tables[0]
	methods := map[string]string{}
	for _, ix := range tb.Indexes {
		methods[ix.Name] = ix.Method
	}
	if got := methods["refresh_tokens_claims_idx"]; got != "gin" {
		t.Errorf("GIN: Method = %q, want %q", got, "gin")
	}
	if got := methods["idx_refresh_tokens_id"]; got != "btree" {
		t.Errorf("BTREE: Method = %q, want %q", got, "btree")
	}
	if got := methods["idx_refresh_tokens_trgm"]; got != "hash" {
		t.Errorf("HASH: Method = %q, want %q", got, "hash")
	}
	if got, ok := methods["idx_refresh_tokens_default"]; !ok || got != "" {
		t.Errorf("default (no USING): Method = %q, want empty string", got)
	}
}
