//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=csrf
//ff:what TestHasCsrf — auth 선언 && csrf 미비활성 시 CSRF 블록 마운트 (BUG-116: bearer 포함)
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestHasCsrf(t *testing.T) {
	cases := []struct {
		name string
		a    prepared.Auth
		want bool
	}{
		{"auth absent", prepared.Auth{Present: false}, false},
		{"present but nil raw", prepared.Auth{Present: true, Raw: nil}, false},
		{
			"raw csrf unset → default enabled",
			prepared.Auth{CsrfRequired: true, Present: true, Mode: "cookie", Raw: &pmanifest.Auth{Csrf: nil}},
			true,
		},
		{
			"csrf explicitly enabled",
			prepared.Auth{CsrfRequired: true, Present: true, Mode: "cookie", Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: true}}},
			true,
		},
		{
			"csrf explicitly disabled",
			prepared.Auth{CsrfRequired: true, Present: true, Mode: "cookie", Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: false}}},
			false,
		},
		{
			// BUG-116 / Phase-B1 — manifest=bearer build (CsrfRequired=false)
			// still mounts the CSRF block: BACKEND_AUTH_MODE can flip the
			// binary to cookie/hybrid at runtime, and the emitted middleware
			// no-ops in bearer mode (csrfRuntimeActive).
			"bearer present → runtime-reachable cookie/hybrid",
			prepared.Auth{CsrfRequired: false, Present: true, Mode: "bearer", Raw: &pmanifest.Auth{Mode: "bearer"}},
			true,
		},
	}
	for _, c := range cases {
		if got := hasCsrf(c.a); got != c.want {
			t.Errorf("%s: hasCsrf = %v, want %v", c.name, got, c.want)
		}
	}
}
