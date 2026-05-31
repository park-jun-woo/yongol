//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestAllTableContents(t *testing.T) {
	files := []sqlFile{
		{content: "CREATE TABLE users (\n  id BIGINT\n);\n"},
		{content: "CREATE TABLE posts (\n  id BIGINT\n);\n"},
	}
	m := allTableContents(files)
	if _, ok := m["users"]; !ok {
		t.Error("missing users key")
	}
	if _, ok := m["posts"]; !ok {
		t.Error("missing posts key")
	}
	if len(m) != 2 {
		t.Errorf("got %d keys, want 2", len(m))
	}
}
