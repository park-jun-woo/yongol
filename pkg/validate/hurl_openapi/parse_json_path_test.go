//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what parseJSONPath — $.a.b[0].c 를 세그먼트 배열로 분해 검증

package hurl_openapi

import "testing"

func TestParseJSONPath(t *testing.T) {
	cases := []struct {
		name string
		path string
		want []string
	}{
		{name: "simple", path: "$.name", want: []string{"name"}},
		{name: "nested", path: "$.user.id", want: []string{"user", "id"}},
		{name: "array_index", path: "$.items[0].name", want: []string{"items", "[0]", "name"}},
		{name: "deep_nested", path: "$.a.b.c.d", want: []string{"a", "b", "c", "d"}},
		{name: "just_dollar", path: "$", want: nil},
		{name: "dollar_dot", path: "$.", want: nil},
		{name: "no_dollar", path: "name", want: []string{"name"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runStringSliceCase(t, parseJSONPath(c.path), c.want)
		})
	}
}
