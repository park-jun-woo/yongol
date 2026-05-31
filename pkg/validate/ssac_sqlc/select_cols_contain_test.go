//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestSelectColsContain(t *testing.T) {
	cols := []string{"id", "email"}
	if !selectColsContain(cols, "email") {
		t.Error("expected email present")
	}
	if selectColsContain(cols, "name") {
		t.Error("expected name absent")
	}
}
