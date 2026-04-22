//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what extractHintComments — 주석 파싱 (-- @rename / @cast / ...)
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractHintComments_RenameAndBackfill(t *testing.T) {
	dir := t.TempDir()
	body := `-- @rename from=members to=users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    -- @rename from=email
    email_address VARCHAR(255) NOT NULL,
    -- @backfill default=false
    email_verified BOOLEAN NOT NULL,
    role TEXT -- @cast using=role::text
);
`
	if err := os.WriteFile(filepath.Join(dir, "users.sql"), []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
	hints, err := ExtractHintCommentsFromDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(hints) < 4 {
		t.Fatalf("expected >=4 hint comments, got %d: %+v", len(hints), hints)
	}
	// Table rename
	foundTableRename := false
	for _, h := range hints {
		if h.Tag == "rename" && h.Args["from"] == "members" && h.Args["to"] == "users" && h.BlockAbove {
			foundTableRename = true
		}
	}
	if !foundTableRename {
		t.Errorf("table rename not found: %+v", hints)
	}
	// Column rename
	foundColRename := false
	for _, h := range hints {
		if h.Tag == "rename" && h.Args["from"] == "email" && h.ColumnCtx == "email_address" {
			foundColRename = true
		}
	}
	if !foundColRename {
		t.Errorf("column rename not found: %+v", hints)
	}
	// Backfill
	foundBackfill := false
	for _, h := range hints {
		if h.Tag == "backfill" && h.Args["default"] == "false" && h.ColumnCtx == "email_verified" {
			foundBackfill = true
		}
	}
	if !foundBackfill {
		t.Errorf("backfill not found: %+v", hints)
	}
}
