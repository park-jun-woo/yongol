//ff:func feature=gen-gogin type=generator control=sequence
//ff:what generateCookieAuthDomain — cookieauth_<ident>.go 1파일 생성 (CookieAuthStrict<Title>, 쿠키 직접 추출)

package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateCookieAuthDomain writes internal/middleware/cookieauth_<ident>.go for
// one cookie/hybrid domain (Phase008 §3a/§3c/§3d). The emitted
// CookieAuthStrict<Title> returns the domain's api_<ident>.StrictMiddlewareFunc
// and reads the token via auth.ExtractAccessFromCookie(ctx) directly — it does
// NOT call the shared extractToken()/authMode(). The domain-suffixed filename
// and func name keep 1-file-1-func and avoid redeclaration across domains.
func generateCookieAuthDomain(mwDir, modulePath, name string) error {
	ident := domainIdent(name)
	title := domainTitle(name)
	apiPkg := "api_" + ident
	apiImport := fmt.Sprintf("%s/internal/%s", modulePath, apiPkg)

	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{Feature: "middleware", Type: "middleware", Control: "sequence", Topic: "auth-check"},
		What: fmt.Sprintf("CookieAuthStrict%s — %s 도메인 세션 토큰 검증 (쿠키 직접 추출)", title, ident),
	})
	body := renderDomainAuthMiddleware(cookieAuthStrictDomainTemplate, apiImport, modulePath, apiPkg, title)
	file := authMwFileName("cookieauth", ident)
	if err := os.WriteFile(filepath.Join(mwDir, file), []byte(header+body), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", file, err)
	}
	return nil
}
