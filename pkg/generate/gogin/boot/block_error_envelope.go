//ff:func feature=gen-gogin type=generator control=sequence topic=error-envelope
//ff:what blockErrorEnvelope — middleware.ErrorEnvelopeMiddleware 등록 (Phase004)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockErrorEnvelope emits r.Use(middleware.ErrorEnvelopeMiddleware()) and
// sets the ExposeInternalError package-level flag from
// BACKEND_ERROR_EXPOSE_INTERNAL_ERROR. Positioned right after blockRequestID
// so a request id is always available for envelope bodies and defensive
// downstream abort handling.
//
// Always active. manifest.backend.error.expose_internal_error provides the
// compile-time default; env flag overrides at runtime (dev convenience).
func blockErrorEnvelope(fs *yongol.Fullstack, modulePath string) MainBlock {
	exposeDefault := resolveExposeInternalError(fs)
	lines := []string{
		fmt.Sprintf(`middleware.ExposeInternalError = envBool("BACKEND_ERROR_EXPOSE_INTERNAL_ERROR", %v)`, exposeDefault),
		`r.Use(middleware.ErrorEnvelopeMiddleware())`,
	}
	return MainBlock{
		Name: "error-envelope",
		Imports: []string{
			fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
		},
		Lines: lines,
	}
}
