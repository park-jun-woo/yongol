//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockRouter — gin.Default + (OTel 활성 시) otelgin 미들웨어 등록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockRouter produces gin engine creation and oapi-codegen route registration.
// Always active.
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
		`"` + modulePath + `/internal/api"`,
	}
	lines := []string{
		`r := gin.Default()`,
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
