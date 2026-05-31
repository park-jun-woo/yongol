//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestWidthToPGCast(t *testing.T) {
	if got := widthToPGCast("int64"); got != "bigint" {
		t.Errorf("widthToPGCast(int64) = %q, want bigint", got)
	}
	if got := widthToPGCast("int32"); got != "int" {
		t.Errorf("widthToPGCast(int32) = %q, want int", got)
	}
	if got := widthToPGCast(""); got != "int" {
		t.Errorf("widthToPGCast(empty) = %q, want int", got)
	}
}
