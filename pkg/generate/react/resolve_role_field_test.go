//ff:func feature=gen-react type=test control=sequence
//ff:what resolveRoleField — nil fs/manifest/auth 의 "" 폴백과 선언된 role_field 반환 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveRoleField(t *testing.T) {
	if got := resolveRoleField(nil); got != "" {
		t.Errorf("nil fs: got %q, want \"\"", got)
	}
	if got := resolveRoleField(&yongol.Fullstack{}); got != "" {
		t.Errorf("nil manifest: got %q, want \"\"", got)
	}
	noAuth := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if got := resolveRoleField(noAuth); got != "" {
		t.Errorf("nil frontend.auth: got %q, want \"\"", got)
	}
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{RoleField: "role"}
	if got := resolveRoleField(fs); got != "role" {
		t.Errorf("declared role_field: got %q, want %q", got, "role")
	}
}
