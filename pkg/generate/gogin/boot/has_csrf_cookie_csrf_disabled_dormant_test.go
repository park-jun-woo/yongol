//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestHasCsrf_CookieCsrfDisabled_Dormant — csrf.enabled=false 이면 hasCsrf false

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestHasCsrf_CookieCsrfDisabled_Dormant(t *testing.T) {
	raw := &pmanifest.Auth{
		Mode: "cookie",
		Csrf: &pmanifest.CsrfConfig{Enabled: false},
	}
	a := prepared.Auth{Present: true, Mode: "cookie", CsrfRequired: true, Raw: raw}
	if hasCsrf(a) {
		t.Fatalf("csrf.enabled=false must report hasCsrf=false (rejected earlier by SEC-201 at validate)")
	}
}
