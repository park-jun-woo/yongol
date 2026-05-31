//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsArchivedTable(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Flags: rule.StringSet{"archived.old_logs": true}})
	if !isArchivedTable(fs, "old_logs") {
		t.Error("old_logs should be archived")
	}
	if isArchivedTable(fs, "users") {
		t.Error("users not archived")
	}
	// nil Ground → false.
	if isArchivedTable(&yongol.Fullstack{}, "old_logs") {
		t.Error("nil Ground → false")
	}
}
