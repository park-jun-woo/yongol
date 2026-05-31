//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"strings"
	"testing"
)

func TestBuildXqs19Advice(t *testing.T) {
	got := buildXqs19Advice("session", "GetUser")
	if !strings.Contains(got, "GetUser") || !strings.Contains(got, "specs/db/queries/session.sql") {
		t.Errorf("advice missing parts: %q", got)
	}
}
