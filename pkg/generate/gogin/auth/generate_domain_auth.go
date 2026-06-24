//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateDomainAuth — 도메인별 strict 미들웨어 파일 생성 (mode 별 bearer/cookie, 공유 헬퍼 미방출)

package auth

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateDomainAuth emits one strict-middleware file per manifest domain
// (Phase008). Each domain's resolved auth_mode (its own override, else the
// backend mode p.Auth.Mode) selects the emitter: bearer →
// bearerauth_<ident>.go, cookie/hybrid → cookieauth_<ident>.go. The shared
// helpers auth_mode.go / extract_token.go are deliberately NOT written here —
// per-domain middleware bypasses the global mode switch (§3c), so they remain
// single-site only.
func generateDomainAuth(fs *yongol.Fullstack, p prepared.State, artifactsDir, modulePath string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return err
	}
	for _, name := range fs.DomainNames() {
		mode := fs.Manifest.Domains[name].ResolvedAuthMode(p.Auth.Mode)
		if err := emitDomainAuthFile(mwDir, modulePath, name, mode); err != nil {
			return err
		}
	}
	return nil
}
