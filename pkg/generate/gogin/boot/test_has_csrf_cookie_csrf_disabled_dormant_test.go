//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestHasCsrf_CookieCsrfDisabled_Dormant — csrf.enabled=false 이면 hasCsrf false

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestHasCsrf_CookieCsrfDisabled_Dormant(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Mode: "cookie",
					Csrf: &pmanifest.CsrfConfig{Enabled: false},
				},
			},
		},
	}
	if hasCsrf(fs) {
		t.Fatalf("csrf.enabled=false must report hasCsrf=false (rejected earlier by SEC-201 at validate)")
	}
}
