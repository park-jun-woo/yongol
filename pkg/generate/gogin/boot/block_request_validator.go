//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockRequestValidator — request_validator 미들웨어 등록 (CORS 이후, Health 이전)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockRequestValidator emits the request-validator bootstrap block. Positioned
// after CORS and before health so /health /ready bypass is handled by the
// middleware's own prefix check and normal API routes undergo OpenAPI
// constraint validation (minLength / pattern / required / enum …).
//
// The middleware constructor returns (gin.HandlerFunc, error) — bootstrap
// failures (malformed openapi.yaml, router build error) log a structured
// error and exit with code 1 instead of panicking. This keeps restart loops
// observable via slog and produces a clear stderr line for operators.
func blockRequestValidator(fs *yongol.Fullstack, modulePath string) MainBlock {
	// Domain mode (BUG-142): there is no single global validator. Each domain's
	// validator is mounted on its own route group in blockRegisterHandlersDomained
	// (appendDomainHandler → group.Use), so emit nothing here.
	if fs.IsDomained() {
		return MainBlock{Name: "request-validator"}
	}
	return MainBlock{
		Name: "request-validator",
		Imports: []string{
			`"fmt"`,
			`"log/slog"`,
			`"os"`,
			fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
		},
		Lines: []string{
			`validator, err := middleware.RequestValidator()`,
			`if err != nil {`,
			`	slog.Error("bootstrap failed", "stage", "request-validator", "err", err)`,
			`	fmt.Fprintf(os.Stderr, "bootstrap failed: %v\n", err)`,
			`	os.Exit(1)`,
			`}`,
			`r.Use(validator)`,
		},
	}
}
