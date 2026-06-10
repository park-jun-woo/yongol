//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=csrf
//ff:what TestCsrfActive — auth 선언 여부 기반 csrf.go 방출 게이트 검증 (BUG-116: bearer 포함)
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
		{"auth-absent", prepared.Auth{Present: false}, false},
		{"present-nil-raw", prepared.Auth{Present: true, Raw: nil}, false},
		{"present-nil-csrf-defaults-on", prepared.Auth{CsrfRequired: true, Present: true, Mode: "cookie", Raw: &pmanifest.Auth{}}, true},
		{"csrf-enabled", prepared.Auth{CsrfRequired: true, Present: true, Mode: "cookie", Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: true}}}, true},
		{"csrf-disabled", prepared.Auth{CsrfRequired: true, Present: true, Mode: "cookie", Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: false}}}, false},
		// BUG-116 / Phase-B1 — manifest=bearer build (CsrfRequired=false) still
		// emits csrf.go: BACKEND_AUTH_MODE can flip the binary to cookie/hybrid
		// at runtime, and the emitted middleware no-ops in bearer mode.
		{"bearer-present-runtime-reachable", prepared.Auth{CsrfRequired: false, Present: true, Mode: "bearer", Raw: &pmanifest.Auth{Mode: "bearer"}}, true},
	}
	for _, c := range cases {
		if got := csrfActive(c.auth); got != c.want {
			t.Errorf("%s: csrfActive = %v, want %v", c.name, got, c.want)
		}
	}
}
