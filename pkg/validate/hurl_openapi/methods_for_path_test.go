//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what methodsForPath — segs 일치하는 route의 HTTP method 정렬 리스트 검증

package hurl_openapi

import (
	"testing"
)

func TestMethodsForPath(t *testing.T) {
	routes := []apiRoute{
		{Path: "/users", Method: "GET", Segments: []string{"users"}},
		{Path: "/users", Method: "POST", Segments: []string{"users"}},
		{Path: "/users/{id}", Method: "DELETE", Segments: []string{"users", ":param"}},
		{Path: "/orders", Method: "GET", Segments: []string{"orders"}},
	}

	cases := []struct {
		name string
		segs []string
		want []string
	}{
		{name: "match_users", segs: []string{"users"}, want: []string{"GET", "POST"}},
		{name: "match_users_param", segs: []string{"users", ":param"}, want: []string{"DELETE"}},
		{name: "match_orders", segs: []string{"orders"}, want: []string{"GET"}},
		{name: "no_match", segs: []string{"products"}, want: nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runStringSliceCase(t, methodsForPath(c.segs, routes), c.want)
		})
	}
}
