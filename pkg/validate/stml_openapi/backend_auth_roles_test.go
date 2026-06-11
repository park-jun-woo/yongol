//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what backendAuthRoles — roles 접근자(nil-safe/미선언/선언) 검증

package stml_openapi

import "testing"

func TestBackendAuthRoles(t *testing.T) {
	if got := backendAuthRoles(nil); got != nil {
		t.Errorf("nil fs: want nil, got %v", got)
	}
	fs := makeFS(nil, nil)
	if got := backendAuthRoles(fs); got != nil {
		t.Errorf("no backend.auth: want nil, got %v", got)
	}
	fs = makeAuthFS(nil, nil, "cookie")
	fs.Manifest.Backend.Auth.Roles = []string{"member", "admin"}
	if got := backendAuthRoles(fs); len(got) != 2 || got[0] != "member" {
		t.Errorf("roles: got %v", got)
	}
}
