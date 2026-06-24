//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what blockRegisterHandlers — strict-server NewStrictHandler + per-op 미들웨어 + RegisterHandlers
package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockRegisterHandlers produces the oapi-codegen strict-server handler
// registration. When bearerAuth middleware is active, it also emits a
// publicOps map literal (collected from the OpenAPI doc) and injects
// BearerAuthStrict into NewStrictHandler's middleware slice. This replaces
// the earlier global r.Use(BearerAuth) pattern so per-op security opt-out
// (OpenAPI `security: []`) is honored.
//
// Generic application-layer rate limiting was retired in favour of
// CDN/WAF/Gateway enforcement (see plans/deprecated/Phase006).
// Phase003 (BUG-089) re-introduced auth-scoped rate limiting: manifest
// backend.rate_limit is mandatory when backend.auth exists, and codegen
// wires RouteRateLimit middleware via blockRateLimit.
func blockRegisterHandlers(fs *yongol.Fullstack, modulePath string) MainBlock {
	// Domain mode (Phase007): mount each domain's handlers on its own route
	// group, all sharing the single srv (Decision B). Single-site keeps the
	// historical single RegisterHandlers below.
	if fs.IsDomained() {
		return blockRegisterHandlersDomained(fs, modulePath)
	}
	imports := []string{fmt.Sprintf(`"%s/internal/api"`, modulePath)}
	var lines []string

	mwFuncs := []string{}
	if hasBearerAuth(fs) {
		imports = append(imports, fmt.Sprintf(`"%s/internal/middleware"`, modulePath))
		publicOps := collectPublicOps(fs.OpenAPIDoc)

		lines = append(lines, "publicOps := map[string]bool{")
		for _, opID := range publicOps {
			lines = append(lines, fmt.Sprintf("\t%q: true,", opID))
		}
		lines = append(lines, "}")
		// Phase003 — secret no longer threaded through. BearerAuthStrict
		// calls auth.VerifyToken, which reads os.Getenv(Config.SecretEnv)
		// set by auth.Configure in blockAuthInit (Phase009).
		mwFuncs = append(mwFuncs, "middleware.BearerAuthStrict(publicOps)")
	}

	lines = append(lines, "strictHandler := api.NewStrictHandler(srv, []api.StrictMiddlewareFunc{")
	for _, fn := range mwFuncs {
		lines = append(lines, "\t"+fn+",")
	}
	lines = append(lines, "})")
	lines = append(lines, "api.RegisterHandlers(r, strictHandler)")

	return MainBlock{
		Name:    "register-handlers",
		Imports: imports,
		Lines:   lines,
	}
}
