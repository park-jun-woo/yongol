//ff:func feature=gen-gogin type=generator control=sequence
//ff:what generateBearerAuthDomain — bearerauth_<ident>.go 1파일 생성 (BearerAuthStrict<Title>, 헤더 인라인 추출)

package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateBearerAuthDomain writes internal/middleware/bearerauth_<ident>.go for
// one bearer domain (Phase008 §3b/§3c/§3d). The emitted BearerAuthStrict<Title>
// returns the domain's api_<ident>.StrictMiddlewareFunc and inlines the
// Authorization header parse — it does NOT call the shared extractToken()/
// authMode(). The domain-suffixed filename and func name prevent overwrite /
// redeclaration when two domains are bearer.
func generateBearerAuthDomain(mwDir, modulePath, name string) error {
	ident := domainIdent(name)
	title := domainTitle(name)
	apiPkg := "api_" + ident
	apiImport := fmt.Sprintf("%s/internal/%s", modulePath, apiPkg)

	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{Feature: "middleware", Type: "middleware", Control: "sequence", Topic: "auth-check"},
		What: fmt.Sprintf("BearerAuthStrict%s — %s 도메인 세션 토큰 검증 (Authorization 헤더 인라인 추출)", title, ident),
	})
	body := renderDomainAuthMiddleware(bearerAuthStrictDomainTemplate, apiImport, modulePath, apiPkg, title)
	file := authMwFileName("bearerauth", ident)
	if err := os.WriteFile(filepath.Join(mwDir, file), []byte(header+body), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", file, err)
	}
	return nil
}
