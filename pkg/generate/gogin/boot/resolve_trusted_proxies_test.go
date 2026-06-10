//ff:func feature=gen-gogin type=test control=sequence topic=trusted-proxy
//ff:what resolveTrustedProxies — manifest.backend.http.trusted_proxies CIDR 목록 추출 (미설정 시 nil)

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveTrustedProxies(t *testing.T) {
	t.Run("NilFullstack", func(t *testing.T) {
		if got := resolveTrustedProxies(nil); got != nil {
			t.Errorf("nil fullstack must resolve to nil, got %v", got)
		}
	})

	t.Run("NoHTTPConfig", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
		if got := resolveTrustedProxies(fs); got != nil {
			t.Errorf("missing backend.http must resolve to nil, got %v", got)
		}
	})

	t.Run("EmptyList", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					HTTP: &pmanifest.HTTPConfig{TrustedProxies: []string{}},
				},
			},
		}
		if got := resolveTrustedProxies(fs); got != nil {
			t.Errorf("empty trusted_proxies must resolve to nil, got %v", got)
		}
	})

	t.Run("DeclaredCIDRs", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					HTTP: &pmanifest.HTTPConfig{
						TrustedProxies: []string{"10.0.0.0/8", "172.16.0.0/12"},
					},
				},
			},
		}
		got := resolveTrustedProxies(fs)
		if len(got) != 2 || got[0] != "10.0.0.0/8" || got[1] != "172.16.0.0/12" {
			t.Errorf("declared CIDRs must pass through, got %v", got)
		}
	})
}
