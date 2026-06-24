//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockDomainCORS — 멀티 도메인 CORS 블록 (AllowOrigins 제거, AllowOriginWithContextFunc + 경로 분기)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockDomainCORS builds the multi-domain CORS MainBlock (Phase008 §2, Decision
// L). CORS stays global but AllowOrigins MUST be cleared — leaving it populated
// would OR-merge with the context func and allow every global origin on
// /api/admin (gin-contrib/cors v1.7.7 isOriginValid checks AllowOrigins first).
// buildCORSConfig therefore sets ONLY AllowOriginWithContextFunc, dispatching
// to the per-path isDomainOriginAllowed helper.
func blockDomainCORS(fs *yongol.Fullstack, c *manifest.CORSConfig, lines []string) MainBlock {
	dispatchFunc, dispatchImports := buildDomainCORSDispatch(fs)
	helperFunc := fmt.Sprintf(`func buildCORSConfig() cors.Config {
	return cors.Config{
		AllowMethods:     envStringList("CORS_ALLOW_METHODS", %s),
		AllowHeaders:     %s,
		ExposeHeaders:    %s,
		AllowCredentials: envBool("CORS_ALLOW_CREDENTIALS", %v),
		MaxAge:           %s,
		AllowOriginWithContextFunc: func(c *gin.Context, origin string) bool {
			return isDomainOriginAllowed(origin, c.Request.URL.Path)
		},
	}
}`, goStringSlice(c.AllowMethods), goStringSlice(c.AllowHeaders),
		goStringSlice(c.ExposeHeaders), c.AllowCredentials,
		durationLiteral(c.MaxAge))

	imports := append([]string{
		`"github.com/gin-contrib/cors"`,
		`"github.com/gin-gonic/gin"`,
	}, dispatchImports...)

	return MainBlock{
		Name:    "cors",
		Imports: imports,
		Lines:   lines,
		Funcs:   []string{helperFunc, dispatchFunc},
	}
}
