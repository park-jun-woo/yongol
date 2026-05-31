//ff:func feature=gen-filefunc type=test control=iteration dimension=1
//ff:what TestResolveFeatureDescription — SSOT→infra→fallback 우선순위 검증
package filefunc

import (
	"testing"
)

func TestResolveFeatureDescription(t *testing.T) {
	cases := []struct {
		name     string
		ssotDesc string
		want     string
	}{
		{"auth", "custom desc", "custom desc"},        // SSOT wins
		{"auth", "", "JWT issuance and verification"}, // infra baseline
		{"totally-unknown", "", fallbackDescription},  // fallback
	}
	for _, c := range cases {
		if got := resolveFeatureDescription(c.name, c.ssotDesc); got != c.want {
			t.Errorf("resolveFeatureDescription(%q,%q) = %q, want %q", c.name, c.ssotDesc, got, c.want)
		}
	}
}
