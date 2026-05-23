//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what findPathMatch — method 무시하고 segs만 일치하는 첫 route index 검증

package hurl_openapi

import "testing"

func TestFindPathMatch(t *testing.T) {
	routes := []apiRoute{
		{Path: "/users", Method: "GET", Segments: []string{"users"}},
		{Path: "/users/{id}", Method: "GET", Segments: []string{"users", ":param"}},
		{Path: "/orders", Method: "POST", Segments: []string{"orders"}},
	}

	cases := []struct {
		name string
		segs []string
		want int
	}{
		{name: "match_users", segs: []string{"users"}, want: 0},
		{name: "match_users_param", segs: []string{"users", ":param"}, want: 1},
		{name: "match_orders", segs: []string{"orders"}, want: 2},
		{name: "no_match", segs: []string{"products"}, want: -1},
		{name: "empty_segs_no_match", segs: nil, want: -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := findPathMatch(c.segs, routes)
			if got != c.want {
				t.Errorf("findPathMatch(...) = %d, want %d", got, c.want)
			}
		})
	}

	t.Run("empty_routes", func(t *testing.T) {
		got := findPathMatch([]string{"users"}, nil)
		if got != -1 {
			t.Errorf("expected -1, got %d", got)
		}
	})
}
