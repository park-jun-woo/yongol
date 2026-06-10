//ff:func feature=gen-react type=generator control=sequence
//ff:what 단일 레이아웃 컴포넌트의 TSX 소스를 생성한다 (data-nav 페이지명 치환 + data-logout 모드별 방출)

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderLayoutTSX generates the full TSX source for a single layout
// component. data-nav values resolve through navLinkPath (static path
// verbatim, page-name reference → resolved route — page-flow Phase010);
// data-logout emits a mode-wired logout button when the project has auth
// (authMode "bearer" or "cookie" — the prepared.AuthFor derivation), and
// is skipped entirely without auth (TM-38 surfaces the dead declaration).
// The layout is the first component class importing the api client and
// session store — the import paths follow the page emitter convention
// (@/lib/api, @/stores/auth).
func renderLayoutTSX(componentName string, layout stml.LayoutSpec, routePatterns map[string]string, authMode string) string {
	emitLogout := layout.Logout != nil && authMode != ""
	var sb strings.Builder

	imports := layoutImports(layout, emitLogout)
	fmt.Fprintf(&sb, "import { %s } from 'react-router-dom'\n", strings.Join(imports, ", "))
	if emitLogout && authMode == "bearer" {
		sb.WriteString("import { useAuthStore } from '@/stores/auth'\n")
	}
	if emitLogout && layout.Logout.OperationID != "" {
		sb.WriteString("import { api } from '@/lib/api'\n")
	}

	fmt.Fprintf(&sb, "\nexport default function %s() {\n", componentName)
	if emitLogout {
		sb.WriteString("  const navigate = useNavigate()\n")
		sb.WriteString(renderLogoutHandler(layout.Logout.OperationID, authMode))
		sb.WriteString("\n")
	}
	sb.WriteString("  return (\n")
	sb.WriteString("    <div>\n")

	renderLayoutNav(&sb, layout, routePatterns, emitLogout)

	if layout.HasOutlet {
		sb.WriteString("      <Outlet />\n")
	}

	sb.WriteString("    </div>\n")
	sb.WriteString("  )\n")
	sb.WriteString("}\n")

	return sb.String()
}
