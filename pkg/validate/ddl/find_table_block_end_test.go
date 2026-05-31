//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestFindTableBlockEnd(t *testing.T) {
	lines := []string{
		"CREATE TABLE foo (",
		"  id BIGINT",
		");",
		"CREATE TABLE bar (",
		"  x TEXT",
	}
	if got := findTableBlockEnd(lines, 0); got != 2 {
		t.Errorf("findTableBlockEnd terminated block = %d, want 2", got)
	}
	// No terminator after start index → returns last index.
	if got := findTableBlockEnd(lines, 3); got != len(lines)-1 {
		t.Errorf("findTableBlockEnd unterminated = %d, want %d", got, len(lines)-1)
	}
}
