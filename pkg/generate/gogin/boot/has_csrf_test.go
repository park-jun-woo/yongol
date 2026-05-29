//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what hasCsrf — prepared.Auth.Mode=cookie|hybrid && csrf.enabled 여부

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
		{"csrf not required (bearer)", prepared.Auth{CsrfRequired: false}, false},
		{"required but auth absent", prepared.Auth{CsrfRequired: true, Present: false}, false},
		{"required but nil raw", prepared.Auth{CsrfRequired: true, Present: true, Raw: nil}, false},
		{
			"required, raw csrf unset → default enabled",
			prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{Csrf: nil}},
			true,
		},
		{
			"required, csrf explicitly enabled",
			prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: true}}},
			true,
		},
		{
			"required, csrf explicitly disabled",
			prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: false}}},
			false,
		},
	}
	for _, c := range cases {
		if got := hasCsrf(c.a); got != c.want {
			t.Errorf("%s: hasCsrf = %v, want %v", c.name, got, c.want)
		}
	}
}
