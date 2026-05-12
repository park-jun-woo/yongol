//ff:func feature=gen-react type=generator control=sequence
//ff:what 단일 레이아웃 컴포넌트의 TSX 소스를 생성한다

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderLayoutTSX generates the full TSX source for a single layout component.
func renderLayoutTSX(componentName string, layout stml.LayoutSpec) string {
	var sb strings.Builder

	imports := layoutImports(layout)
	fmt.Fprintf(&sb, "import { %s } from 'react-router-dom'\n", strings.Join(imports, ", "))

	fmt.Fprintf(&sb, "\nexport default function %s() {\n", componentName)
	sb.WriteString("  return (\n")
	sb.WriteString("    <div>\n")

	if len(layout.NavItems) > 0 {
		sb.WriteString("      <nav>\n")
		for _, item := range layout.NavItems {
			fmt.Fprintf(&sb, "        <Link to=\"%s\">%s</Link>\n", item.Path, item.Label)
		}
		sb.WriteString("      </nav>\n")
	}

	if layout.HasOutlet {
		sb.WriteString("      <Outlet />\n")
	}

	sb.WriteString("    </div>\n")
	sb.WriteString("  )\n")
	sb.WriteString("}\n")

	return sb.String()
}
