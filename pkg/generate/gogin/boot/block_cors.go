//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockCORS — gin-contrib/cors 미들웨어 등록 (manifest + env 기반)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockCORS emits a cors.New(cors.Config{...}) middleware registration based
// on manifest.backend.cors. When CORS is disabled or unconfigured, the block
// is inert (empty Lines). envStringList / envBool (declared by
// blockEnvHelpers) provide runtime overrides for origins, methods, and
// credentials.
//
// Positioned after blockRouter and before blockHealth so health/ready
// probes also benefit from CORS headers (browser status dashboards).
func blockCORS(fs *yongol.Fullstack) MainBlock {
	if !corsEnabled(fs) {
		return MainBlock{Name: "cors"}
	}
	c := fs.Manifest.Backend.CORS

	lines := []string{
		`r.Use(cors.New(buildCORSConfig()))`,
	}

	// Multi-domain (Phase008 §2, Decision L): origin validation becomes
	// path-aware. Delegated to blockDomainCORS so blockCORS stays small.
	if fs.IsDomained() {
		return blockDomainCORS(fs, c, lines)
	}

	helperFunc := fmt.Sprintf(`func buildCORSConfig() cors.Config {
	cfg := cors.Config{
		AllowMethods:     envStringList("CORS_ALLOW_METHODS", %s),
		AllowHeaders:     %s,
		ExposeHeaders:    %s,
		AllowCredentials: envBool("CORS_ALLOW_CREDENTIALS", %v),
		MaxAge:           %s,
	}
	cfg.AllowOrigins = envStringList("CORS_ALLOW_ORIGINS", %s)
	return cfg
}`, goStringSlice(c.AllowMethods), goStringSlice(c.AllowHeaders),
		goStringSlice(c.ExposeHeaders), c.AllowCredentials,
		durationLiteral(c.MaxAge), goStringSlice(c.AllowOrigins))

	return MainBlock{
		Name:    "cors",
		Imports: []string{`"github.com/gin-contrib/cors"`},
		Lines:   lines,
		Funcs:   []string{helperFunc},
	}
}
