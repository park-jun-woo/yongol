//ff:func feature=gen-gogin type=generator control=sequence topic=trusted-proxy
//ff:what blockRouter — gin.Default + SetTrustedProxies 안전 기본값 + (OTel 활성 시) otelgin 미들웨어 등록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockRouter produces gin engine creation and oapi-codegen route registration.
// Always active.
//
// Immediately after gin.Default() the block emits an explicit
// r.SetTrustedProxies(...) call (BUG-117). gin's default is to trust ALL
// proxies, which lets any client spoof c.ClientIP() via X-Forwarded-For —
// bypassing IP-keyed rate limiters, IP logging, and IP policies. The
// emitted default is nil (trust no proxy); manifest
// backend.http.trusted_proxies declares trusted CIDR ranges for reverse
// proxy deployments, and BACKEND_HTTP_TRUSTED_PROXIES (comma-separated
// CIDRs) overrides at runtime (env > manifest > default nil). Invalid
// CIDRs make SetTrustedProxies return an error → fail-fast at bootstrap.
//
// When OpenTelemetry tracing is enabled (Phase009), otelgin.Middleware is
// registered FIRST (before request-id / error-envelope / cors / prometheus)
// so the HTTP server span covers every downstream middleware timing. The
// service name attribute comes from the same otelServiceName computed in
// the OTel init block so all spans share a consistent `service.name`
// dimension.
func blockRouter(fs *yongol.Fullstack, modulePath string) MainBlock {
	imports := []string{
		`"github.com/gin-gonic/gin"`,
		`"log/slog"`,
		`"os"`,
		`"` + modulePath + `/internal/api"`,
	}
	trustedDefault := "nil"
	if trusted := resolveTrustedProxies(fs); len(trusted) > 0 {
		trustedDefault = goStringSlice(trusted)
	}
	lines := []string{
		`r := gin.Default()`,
		fmt.Sprintf(`trustedProxies := envStringList("BACKEND_HTTP_TRUSTED_PROXIES", %s)`, trustedDefault),
		`if err := r.SetTrustedProxies(trustedProxies); err != nil {`,
		`	slog.Error("invalid trusted_proxies CIDR", "error", err)`,
		`	os.Exit(1)`,
		`}`,
	}
	if hasOtel(fs) {
		imports = append(imports,
			`"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"`,
		)
		lines = append(lines,
			fmt.Sprintf(`r.Use(otelgin.Middleware(otelServiceName))`),
		)
	}
	return MainBlock{
		Name:    "router",
		Imports: imports,
		Lines:   lines,
	}
}
