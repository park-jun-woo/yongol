//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestFormatReturningReason(t *testing.T) {
	if got := formatReturningReason(ShapeFull, "*"); got != "RETURNING * → model" {
		t.Errorf("full star: %q", got)
	}
	if got := formatReturningReason(ShapeFull, "id, email"); got != "RETURNING <full column list> → model" {
		t.Errorf("full list: %q", got)
	}
	if got := formatReturningReason(ShapePartial, "id"); got != "RETURNING id → partial Row" {
		t.Errorf("partial: %q", got)
	}
	if got := formatReturningReason(ShapeNone, ""); got != "no RETURNING → model" {
		t.Errorf("none: %q", got)
	}
}
