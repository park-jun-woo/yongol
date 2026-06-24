//ff:func feature=gen-gogin type=test control=sequence topic=dos-guard
//ff:what buildOperationRouteIndex — 도메인 모드: route_prefix 가 붙은 route 키 매핑 검증
package boot

import "testing"

func TestBuildOperationRouteIndex_DomainedPrefixesKeys(t *testing.T) {
	idx := buildOperationRouteIndex(domainedFS(nil))
	// public domain (route_prefix /api): /login → POST /api/login
	if idx["Login"] != "POST /api/login" {
		t.Errorf("Login = %q, want POST /api/login", idx["Login"])
	}
	// admin domain (route_prefix /api/admin): /users → GET /api/admin/users
	if idx["AdminListUsers"] != "GET /api/admin/users" {
		t.Errorf("AdminListUsers = %q, want GET /api/admin/users", idx["AdminListUsers"])
	}
	if len(idx) != 2 {
		t.Errorf("index size = %d, want 2", len(idx))
	}
}
