//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=security-headers
//ff:what hasSecurityHeaders — manifest.backend.security_headers.enabled (기본 true) 여부
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasSecurityHeaders(t *testing.T) {
	tru := true
	fals := false
	cases := []struct {
		name string
		fs   *yongol.Fullstack
		want bool
	}{
		{"nil fs defaults true", nil, true},
		{"nil manifest defaults true", &yongol.Fullstack{}, true},
		{"no block defaults true", fsWithSH(nil), true},
		{"nil enabled defaults true", fsWithSH(&pmanifest.SecurityHeadersConfig{Enabled: nil}), true},
		{"explicitly enabled", fsWithSH(&pmanifest.SecurityHeadersConfig{Enabled: &tru}), true},
		{"explicitly disabled", fsWithSH(&pmanifest.SecurityHeadersConfig{Enabled: &fals}), false},
	}
	for _, c := range cases {
		if got := hasSecurityHeaders(c.fs); got != c.want {
			t.Errorf("%s: hasSecurityHeaders = %v, want %v", c.name, got, c.want)
		}
	}
}
