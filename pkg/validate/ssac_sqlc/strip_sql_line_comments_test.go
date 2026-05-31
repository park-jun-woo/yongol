//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestStripSQLLineComments(t *testing.T) {
	in := "SELECT id -- the id\nFROM users; -- partial\n"
	got := stripSQLLineComments(in)
	want := "SELECT id \nFROM users; \n"
	if got != want {
		t.Errorf("stripSQLLineComments = %q, want %q", got, want)
	}
	// No comments → unchanged.
	if got := stripSQLLineComments("SELECT 1"); got != "SELECT 1" {
		t.Errorf("no-comment passthrough = %q", got)
	}
}
