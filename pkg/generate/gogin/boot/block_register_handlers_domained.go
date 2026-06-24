//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what blockRegisterHandlersDomained — 도메인별 r.Group + NewStrictHandler + RegisterHandlers (공유 srv)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockRegisterHandlersDomained emits, for each manifest domain, a route group
// mounted at the domain's route_prefix and that domain's oapi-codegen
// strict-server registration. All domains share the single srv (one
// *service.Server implements every domain's StrictServerInterface — Decision B,
// operationIds globally unique under XDO-90). Each domain imports its own
// internal/api_<ident> package and, when bearer auth is active, builds a
// per-domain publicOps map from that domain's OpenAPI doc so per-op security
// opt-out is honored. Per-domain auth-mode divergence is refined in Phase008.
func blockRegisterHandlersDomained(fs *yongol.Fullstack, modulePath string) MainBlock {
	bearer := hasBearerAuth(fs)
	var imports []string
	var lines []string
	if bearer {
		imports = append(imports, fmt.Sprintf(`"%s/internal/middleware"`, modulePath))
	}
	for _, name := range fs.DomainNames() {
		imports, lines = appendDomainHandler(imports, lines, fs, name, modulePath, bearer)
	}
	return MainBlock{
		Name:    "register-handlers",
		Imports: imports,
		Lines:   lines,
	}
}
