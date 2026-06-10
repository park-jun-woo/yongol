//ff:func feature=gen-gogin type=test control=sequence topic=trusted-proxy
//ff:what blockRouter — manifest trusted_proxies CIDR 방출 + env 오버라이드(BACKEND_HTTP_TRUSTED_PROXIES) 우선 회귀

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRouter_TrustedProxiesManifestAndEnvOverride(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				HTTP: &pmanifest.HTTPConfig{
					TrustedProxies: []string{"10.0.0.0/8", "172.16.0.0/12"},
				},
			},
		},
	}
	block := blockRouter(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")

	// Manifest CIDRs are emitted as the runtime DEFAULT argument of
	// envStringList — so BACKEND_HTTP_TRUSTED_PROXIES (comma-separated
	// CIDRs) still wins at runtime: env > manifest > nil.
	want := `trustedProxies := envStringList("BACKEND_HTTP_TRUSTED_PROXIES", []string{"10.0.0.0/8", "172.16.0.0/12"})`
	if !strings.Contains(body, want) {
		t.Errorf("manifest trusted_proxies must be emitted as envStringList default, got:\n%s", body)
	}
	if !strings.Contains(body, "r.SetTrustedProxies(trustedProxies)") {
		t.Errorf("must emit SetTrustedProxies call, got:\n%s", body)
	}
}
