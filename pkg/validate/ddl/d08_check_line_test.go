//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"strings"
	"testing"
)

func TestD08CheckLine(t *testing.T) {
	f := sqlFile{path: "schema.sql"}
	blk := tableBlock{tableName: "users", startLine: 1}

	// SERIAL column → diagnostic produced.
	d := d08CheckLine(f, blk, "  id SERIAL,", 1)
	if d == nil {
		t.Fatal("expected diagnostic for SERIAL column")
	}
	if !strings.Contains(d.Message, "[D-8]") || !strings.Contains(d.Message, "users.id") {
		t.Errorf("unexpected message: %q", d.Message)
	}
	if !strings.Contains(d.Advice, "IDENTITY") {
		t.Errorf("advice should mention IDENTITY: %q", d.Advice)
	}
	if d.Line != blk.startLine+1 {
		t.Errorf("line = %d, want %d", d.Line, blk.startLine+1)
	}

	// BIGSERIAL also flagged.
	if got := d08CheckLine(f, blk, "n BIGSERIAL", 0); got == nil {
		t.Error("expected diagnostic for BIGSERIAL")
	}

	// Non-serial column → nil.
	if got := d08CheckLine(f, blk, "  id BIGINT,", 0); got != nil {
		t.Errorf("expected nil for BIGINT, got %+v", got)
	}
	// Skippable line → nil.
	if got := d08CheckLine(f, blk, ");", 0); got != nil {
		t.Error("expected nil for skippable line")
	}
	// Line without a column-name match → nil.
	if got := d08CheckLine(f, blk, "123 SERIAL", 0); got != nil {
		t.Error("expected nil when no leading identifier")
	}
	// Single token (no type) → nil.
	if got := d08CheckLine(f, blk, "id", 0); got != nil {
		t.Error("expected nil for single-token line")
	}
}
