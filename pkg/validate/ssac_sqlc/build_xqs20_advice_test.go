//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"strings"
	"testing"
)

func TestBuildXqs20Advice(t *testing.T) {
	partial := buildXqs20Advice("get", "Model", "GetRow", ShapePartial)
	if !strings.Contains(partial, "expand RETURNING") {
		t.Errorf("partial advice: %q", partial)
	}
	full := buildXqs20Advice("get", "Row", "Model", ShapeFull)
	if !strings.Contains(full, "restrict RETURNING") {
		t.Errorf("full advice: %q", full)
	}
	none := buildXqs20Advice("get", "Row", "Model", ShapeNone)
	if !strings.Contains(none, "@get Model") {
		t.Errorf("none advice: %q", none)
	}
}
