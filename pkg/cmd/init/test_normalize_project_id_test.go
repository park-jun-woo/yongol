//ff:type feature=cli-init type=test
//ff:what test — NormalizeProjectID table-driven cases

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

func TestValidateProjectID(t *testing.T) {
	ok := []string{"Zenflow", "zen_flow", "MyApp", "a", "A1", "my_app2"}
	for _, s := range ok {
		if err := ValidateProjectID(s); err != nil {
			t.Errorf("ValidateProjectID(%q) unexpected error: %v", s, err)
		}
	}
	bad := []string{"", "1abc", "my-app", "my.app", "my/app", "my app"}
	for _, s := range bad {
		if err := ValidateProjectID(s); err == nil {
			t.Errorf("ValidateProjectID(%q) want error, got nil", s)
		}
	}
}
