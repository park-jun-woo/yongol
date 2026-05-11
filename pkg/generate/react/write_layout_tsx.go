//ff:func feature=gen-react type=generator control=sequence
//ff:what LayoutSpec -> React layout TSX 파일 생성 (Link + Outlet)

package react

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// writeLayoutTSX emits a single layout TSX file from a LayoutSpec.
// The file is written to <srcDir>/layouts/<PascalName>Layout.tsx.
//
// Conversion rules:
//   - NavItem → <Link to="path">Label</Link>
//   - HasOutlet → <Outlet />
//   - Layout name → PascalCase + "Layout" suffix (e.g. "app" → "AppLayout")
func writeLayoutTSX(srcDir string, layout stml.LayoutSpec) error {
	layoutsDir := filepath.Join(srcDir, "layouts")
	if err := os.MkdirAll(layoutsDir, 0o755); err != nil {
		return err
	}

	componentName := layoutComponentName(layout.Name)
	src := renderLayoutTSX(componentName, layout)
	return os.WriteFile(filepath.Join(layoutsDir, componentName+".tsx"), []byte(src), 0o644)
}

// writeLayoutsTSX emits layout TSX files for all provided LayoutSpecs.
func writeLayoutsTSX(srcDir string, layouts []stml.LayoutSpec) error {
	for _, l := range layouts {
		if err := writeLayoutTSX(srcDir, l); err != nil {
			return err
		}
	}
	return nil
}

// layoutComponentName converts a layout name to PascalCase + "Layout" suffix.
// e.g. "app" → "AppLayout", "auth" → "AuthLayout", "main-nav" → "MainNavLayout"
func layoutComponentName(name string) string {
	return kebabToPascal(name) + "Layout"
}

// renderLayoutTSX generates the full TSX source for a single layout component.
func renderLayoutTSX(componentName string, layout stml.LayoutSpec) string {
	var sb strings.Builder

	// Imports
	imports := layoutImports(layout)
	fmt.Fprintf(&sb, "import { %s } from 'react-router-dom'\n", strings.Join(imports, ", "))

	// Component
	fmt.Fprintf(&sb, "\nexport default function %s() {\n", componentName)
	sb.WriteString("  return (\n")
	sb.WriteString("    <div>\n")

	// Nav section (only if NavItems exist)
	if len(layout.NavItems) > 0 {
		sb.WriteString("      <nav>\n")
		for _, item := range layout.NavItems {
			fmt.Fprintf(&sb, "        <Link to=\"%s\">%s</Link>\n", item.Path, item.Label)
		}
		sb.WriteString("      </nav>\n")
	}

	// Outlet
	if layout.HasOutlet {
		sb.WriteString("      <Outlet />\n")
	}

	sb.WriteString("    </div>\n")
	sb.WriteString("  )\n")
	sb.WriteString("}\n")

	return sb.String()
}

// layoutImports returns the react-router-dom named imports needed for a layout.
func layoutImports(layout stml.LayoutSpec) []string {
	var imports []string
	if len(layout.NavItems) > 0 {
		imports = append(imports, "Link")
	}
	if layout.HasOutlet {
		imports = append(imports, "Outlet")
	}
	return imports
}
