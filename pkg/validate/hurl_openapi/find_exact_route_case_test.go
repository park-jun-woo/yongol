//ff:func feature=validate type=test-helper control=sequence topic=hurl-openapi
//ff:what runFindExactRoute — TestFindExactRoute table-driven 개별 케이스 검증

package hurl_openapi

import "testing"

var findExactRouteFixtureRoutes = []apiRoute{
	{Path: "/users", Method: "GET", Segments: []string{"users"}},
	{Path: "/users", Method: "POST", Segments: []string{"users"}},
	{Path: "/users/{id}", Method: "GET", Segments: []string{"users", ":param"}},
	{Path: "/orders", Method: "GET", Segments: []string{"orders"}},
}

func runFindExactRoute(t *testing.T, c TestFindExactRouteCase) {
	t.Helper()
	var r []apiRoute
	if c.name != "empty_routes" {
		r = findExactRouteFixtureRoutes
	}
	got := findExactRoute(c.segs, c.method, r)
	if c.wantNil {
		if got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
		return
	}
	if got == nil {
		t.Fatal("expected non-nil route, got nil")
	}
	if got.Path != c.wantPath {
		t.Errorf("path = %q, want %q", got.Path, c.wantPath)
	}
}
