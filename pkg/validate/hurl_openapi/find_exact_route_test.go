//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what findExactRoute — segs/method 일치하는 첫 route 반환 검증

package hurl_openapi

import "testing"

func TestFindExactRoute(t *testing.T) {
	cases := []TestFindExactRouteCase{
		{
			name:     "exact_match_GET_users",
			segs:     []string{"users"},
			method:   "GET",
			wantPath: "/users",
		},
		{
			name:     "exact_match_POST_users",
			segs:     []string{"users"},
			method:   "POST",
			wantPath: "/users",
		},
		{
			name:     "parameterized_match",
			segs:     []string{"users", ":param"},
			method:   "GET",
			wantPath: "/users/{id}",
		},
		{
			name:    "no_method_match",
			segs:    []string{"users"},
			method:  "DELETE",
			wantNil: true,
		},
		{
			name:    "no_path_match",
			segs:    []string{"products"},
			method:  "GET",
			wantNil: true,
		},
		{
			name:     "case_insensitive_method",
			segs:     []string{"orders"},
			method:   "get",
			wantPath: "/orders",
		},
		{
			name:    "empty_routes",
			segs:    []string{"users"},
			method:  "GET",
			wantNil: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runFindExactRoute(t, c)
		})
	}
}
