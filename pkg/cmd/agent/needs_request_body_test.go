//ff:func feature=agent type=test control=sequence
//ff:what TestNeedsRequestBody — post/put/patch 만 request body 필요로 판별 검증

package agent

import "testing"

func TestNeedsRequestBody(t *testing.T) {
	cases := map[string]bool{
		"post":   true,
		"put":    true,
		"patch":  true,
		"get":    false,
		"delete": false,
		"":       false,
	}
	for method, want := range cases {
		if got := needsRequestBody(method); got != want {
			t.Errorf("needsRequestBody(%q) = %v, want %v", method, got, want)
		}
	}
}
