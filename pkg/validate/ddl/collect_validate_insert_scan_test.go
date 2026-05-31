//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestCollectValidateInsertScan(t *testing.T) {
	lines := []string{
		"INSERT INTO users VALUES (",
		"  0, 'sys'",
		");",
		"SELECT 1;",
	}
	r, next := collectValidateInsertScan(lines, 0, "users", true)
	if r.Table != "users" {
		t.Errorf("table = %q, want users", r.Table)
	}
	if !r.Annotated {
		t.Error("expected annotated true")
	}
	if r.StartLine != 1 {
		t.Errorf("startLine = %d, want 1", r.StartLine)
	}
	if next != 3 {
		t.Errorf("next index = %d, want 3", next)
	}
}
