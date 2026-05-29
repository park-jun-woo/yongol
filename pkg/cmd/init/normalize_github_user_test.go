//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestNormalizeGitHubUser — table-driven normalizeGitHubUser cases

package cliinit

import "testing"

func TestNormalizeGitHubUser(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Alice", "alice"},
		{"Park Jun Woo", "parkjunwoo"},
		{"park-jun-woo", "park-jun-woo"},
		{"alice.bob", "alicebob"},
		{"alice@corp", "alicecorp"},
	}
	for _, tc := range cases {
		got := normalizeGitHubUser(tc.in)
		if got != tc.want {
			t.Errorf("normalizeGitHubUser(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
