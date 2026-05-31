//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestCsrfActive(t *testing.T) {
	cases := []struct {
		name string
		auth prepared.Auth
		want bool
	}{
		{"not-required", prepared.Auth{CsrfRequired: false}, false},
		{"required-but-absent", prepared.Auth{CsrfRequired: true, Present: false}, false},
		{"required-nil-raw", prepared.Auth{CsrfRequired: true, Present: true, Raw: nil}, false},
		{"required-nil-csrf-defaults-on", prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}, true},
		{"required-csrf-enabled", prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: true}}}, true},
		{"required-csrf-disabled", prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: false}}}, false},
	}
	for _, c := range cases {
		if got := csrfActive(c.auth); got != c.want {
			t.Errorf("%s: csrfActive = %v, want %v", c.name, got, c.want)
		}
	}
}
