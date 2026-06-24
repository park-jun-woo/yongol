//ff:func feature=gen-gogin type=generator control=selection
//ff:what emitDomainAuthFile — 도메인 mode 에 따라 bearer/cookie strict 미들웨어 파일 1개 방출

package auth

import "fmt"

// emitDomainAuthFile writes one domain's strict-middleware file by mode: cookie
// or hybrid → cookieauth_<ident>.go (CookieAuthStrict<Title>), anything else →
// bearerauth_<ident>.go (BearerAuthStrict<Title>). Split out of
// generateDomainAuth to keep the loop body shallow (filefunc Q1).
func emitDomainAuthFile(mwDir, modulePath, name, mode string) error {
	switch mode {
	case "cookie", "hybrid":
		if err := generateCookieAuthDomain(mwDir, modulePath, name); err != nil {
			return fmt.Errorf("cookie auth %s: %w", name, err)
		}
	default: // bearer
		if err := generateBearerAuthDomain(mwDir, modulePath, name); err != nil {
			return fmt.Errorf("bearer auth %s: %w", name, err)
		}
	}
	return nil
}
