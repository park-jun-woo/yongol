//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderCorsBlock — CORS enableCors() 옵션 포함 렌더링

package boot

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderCorsBlock writes app.enableCors() with origin and credentials options
// when available from the BootPlan CORS config.
func renderCorsBlock(b *strings.Builder, plan *ir.BootPlan) {
	origins, credentials := extractCORSConfig(plan)
	if len(origins) == 0 {
		b.WriteString("  app.enableCors({ credentials: true });\n\n")
		return
	}
	b.WriteString("  app.enableCors({\n")
	b.WriteString("    origin: [\n")
	for _, o := range origins {
		b.WriteString(fmt.Sprintf("      '%s',\n", o))
	}
	b.WriteString("    ],\n")
	if credentials {
		b.WriteString("    credentials: true,\n")
	}
	b.WriteString("  });\n\n")
}
