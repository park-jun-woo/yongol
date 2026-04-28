//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestNormalizeProjectID — table-driven NormalizeProjectID cases

package cliinit

import "testing"

func TestNormalizeProjectID(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Zenflow", "zenflow"},
		{"zen_flow", "zen_flow"},
		{"ZenFlow", "zen_flow"},
		{"MyApp", "my_app"},
		{"APIGateway", "api_gateway"},
		{"my_app", "my_app"},
		{"MyURL2Path", "my_url2_path"},
		{"_Foo_", "foo"},
		{"Foo__Bar", "foo_bar"},
	}
	for _, tc := range cases {
		got := NormalizeProjectID(tc.in)
		if got != tc.want {
			t.Errorf("NormalizeProjectID(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
