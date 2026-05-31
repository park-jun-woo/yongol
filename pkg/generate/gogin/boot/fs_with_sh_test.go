//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what hasSecurityHeaders — manifest.backend.security_headers.enabled (기본 true) 여부
package boot

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithSH(sh *pmanifest.SecurityHeadersConfig) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{SecurityHeaders: sh},
	}}
}
