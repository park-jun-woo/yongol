//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestExtractTableBlocks(t *testing.T) {
	f := sqlFile{
		path:    "schema.sql",
		name:    "schema.sql",
		content: "CREATE TABLE users (\n  id BIGINT\n);\nCREATE TABLE IF NOT EXISTS posts (\n  id BIGINT\n);\n",
	}
	blocks := extractTableBlocks(f)
	if len(blocks) != 2 {
		t.Fatalf("got %d blocks, want 2", len(blocks))
	}
	if blocks[0].tableName != "users" {
		t.Errorf("block0 name = %q, want users", blocks[0].tableName)
	}
	if blocks[1].tableName != "posts" {
		t.Errorf("block1 name = %q, want posts", blocks[1].tableName)
	}
	if blocks[0].startLine != 1 {
		t.Errorf("block0 startLine = %d, want 1", blocks[0].startLine)
	}

	// Content with no CREATE TABLE → no blocks.
	if got := extractTableBlocks(sqlFile{content: "SELECT 1;"}); got != nil {
		t.Errorf("expected nil blocks for non-DDL, got %v", got)
	}
}
