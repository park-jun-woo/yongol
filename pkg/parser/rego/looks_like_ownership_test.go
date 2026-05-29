//ff:func feature=policy type=test control=iteration dimension=1
//ff:what looksLikeOwnership — `@ownership` 주석 패턴 감지

package rego

import "testing"

func TestLooksLikeOwnership(t *testing.T) {
	cases := map[string]bool{
		"# @ownership gig: gigs.client_id": true,
		"#@ownership":                      true,
		"# @ownership":                     true,
		"# normal comment":                 false,
		"allow {":                          false,
		"":                                 false,
	}
	for in, want := range cases {
		if got := looksLikeOwnership(in); got != want {
			t.Errorf("looksLikeOwnership(%q) = %v, want %v", in, got, want)
		}
	}
}
