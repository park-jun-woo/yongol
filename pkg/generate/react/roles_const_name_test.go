//ff:func feature=gen-react type=test control=sequence
//ff:what rolesConstName — 단일/복수/비식별자 살균 상수명 검증

package react

import "testing"

func TestRolesConstName(t *testing.T) {
	if got := rolesConstName([]string{"admin"}); got != "ROLES_admin" {
		t.Errorf("single: got %q", got)
	}
	if got := rolesConstName([]string{"admin", "manager"}); got != "ROLES_admin_manager" {
		t.Errorf("multi: got %q", got)
	}
	if got := rolesConstName([]string{"super-admin"}); got != "ROLES_super_admin" {
		t.Errorf("sanitize: got %q", got)
	}
}
