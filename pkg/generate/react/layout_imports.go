//ff:func feature=gen-react type=util control=sequence
//ff:what 레이아웃에 필요한 react-router-dom named import 목록을 반환한다

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

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
