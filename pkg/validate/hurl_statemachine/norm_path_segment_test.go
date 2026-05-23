//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-statemachine
//ff:what normPathSegment — 한 세그먼트 정규화 검증

package hurl_statemachine

import "testing"

func TestNormPathSegment(t *testing.T) {
	cases := []struct {
		name string
		p    string
		want string
	}{
		{name: "openapi_var", p: "{id}", want: ":param"},
		{name: "hurl_var", p: "{{userId}}", want: ":param"},
		{name: "numeric", p: "123", want: ":param"},
		{name: "literal", p: "users", want: "users"},
		{name: "mixed_text", p: "abc123", want: "abc123"},
		{name: "nil_openapi_re_hurl_var", p: "{{x}}", want: ":param"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := normPathSegment(c.p, reOpenAPIVarKey)
			if got != c.want {
				t.Errorf("normPathSegment(%q) = %q, want %q", c.p, got, c.want)
			}
		})
	}

	t.Run("hurl_var_embedded_not_openapi", func(t *testing.T) {
		// "prefix{{id}}suffix" matches hurl var regex (contains {{...}})
		// but not openapi regex (not ^{...}$)
		got := normPathSegment("prefix{{id}}suffix", reOpenAPIVarKey)
		if got != ":param" {
			t.Errorf("got %q, want :param", got)
		}
	})

	t.Run("nil_openapi_re_falls_through", func(t *testing.T) {
		got := normPathSegment("{id}", nil)
		// With nil openapi regex, {id} is not matched as openapi var,
		// but it's also not a hurl var or numeric, so it stays as-is
		if got != "{id}" {
			t.Errorf("got %q, want {id}", got)
		}
	})
}
