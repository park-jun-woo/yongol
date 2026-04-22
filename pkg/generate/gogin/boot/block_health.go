//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockHealth — /health (liveness) + /ready (readiness) 등록

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// blockHealth registers /health (liveness) and /ready (readiness) endpoints.
// /health always returns 200 (process liveness).
// /ready performs DB PingContext when DDL is present; otherwise behaves like
// /health. Both bypass BearerAuth because they are registered directly on the
// gin engine rather than the oapi-codegen strict handler.
func blockHealth(fs *yongol.Fullstack) MainBlock {
	lines := []string{
		`r.GET("/health", func(c *gin.Context) {`,
		`	c.JSON(200, gin.H{"status": "ok"})`,
		`})`,
	}
	withDB := hasDDL(fs)
	lines = append(lines, readyHandlerLines(withDB)...)
	funcs := []string{}
	if withDB {
		funcs = append(funcs, healthHelperReadyWithDB())
	}
	return MainBlock{
		Name:    "health",
		Imports: []string{`"context"`, `"database/sql"`, `"time"`},
		Lines:   lines,
		Funcs:   funcs,
	}
}
