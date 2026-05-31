//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestScanInsertsWithAnnotations(t *testing.T) {
	content := "" +
		"CREATE TABLE users (id BIGINT);\n" +
		"-- @sentinel\n" +
		"INSERT INTO users VALUES (0, 'system');\n" +
		"INSERT INTO posts VALUES (1, 'hi');\n"
	got := scanInsertsWithAnnotations(content)
	if len(got) != 2 {
		t.Fatalf("got %d inserts, want 2", len(got))
	}
	if got[0].Table != "users" || !got[0].Annotated {
		t.Errorf("first insert = %+v, want users annotated", got[0])
	}
	if got[1].Table != "posts" || got[1].Annotated {
		t.Errorf("second insert = %+v, want posts not annotated", got[1])
	}
}
