//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsPkgModelTable(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Flags: rule.StringSet{"pkgModel.sessions": true}})
	if !isPkgModelTable(fs, "sessions") {
		t.Error("sessions should be a pkg model table")
	}
	if isPkgModelTable(fs, "users") {
		t.Error("users not a pkg model table")
	}
	if isPkgModelTable(&yongol.Fullstack{}, "sessions") {
		t.Error("nil Ground → false")
	}
}
