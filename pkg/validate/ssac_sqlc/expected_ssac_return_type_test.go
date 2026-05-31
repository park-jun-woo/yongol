//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestExpectedSsacReturnType(t *testing.T) {
	if got := expectedSsacReturnType(ShapePartial, "User", "GetUser"); got != "GetUserRow" {
		t.Errorf("partial: %q, want GetUserRow", got)
	}
	if got := expectedSsacReturnType(ShapeFull, "User", "GetUser"); got != "User" {
		t.Errorf("full: %q, want User", got)
	}
	if got := expectedSsacReturnType(ShapeNone, "User", "GetUser"); got != "User" {
		t.Errorf("none: %q, want User", got)
	}
	if got := expectedSsacReturnType(ShapeFull, "", "GetUser"); got != "" {
		t.Errorf("empty model: %q, want empty", got)
	}
}
