//ff:func feature=validate type=test-helper control=sequence topic=manifest-auth
//ff:what fsWithAccessTTL — access_token_ttl 값만 지정한 Fullstack 생성

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// fsWithAccessTTL returns a minimal Fullstack with only auth.access_token_ttl set.
// Used by SEC-402 tests to exercise upper-bound boundary cases.
func fsWithAccessTTL(ttl string) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{AccessTokenTTL: ttl},
			},
		},
	}
}
