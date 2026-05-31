//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestReturningCoversAllColumns(t *testing.T) {
	tbl := ddlTable("id", "email")
	full := map[string]bool{"id": true, "email": true}
	if !returningCoversAllColumns(full, tbl.Columns) {
		t.Error("expected full coverage")
	}
	partial := map[string]bool{"id": true}
	if returningCoversAllColumns(partial, tbl.Columns) {
		t.Error("expected partial NOT to cover all")
	}
}
