//ff:func feature=gen-react type=util control=sequence
//ff:what 레이아웃에 필요한 react-router-dom named import 목록을 반환한다 (logout 방출 시 useNavigate 포함)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// layoutImports returns the react-router-dom named imports needed for a
// layout. emitLogout reports whether the layout emits the data-logout
// button (declared AND auth present — page-flow Phase010), which needs
// useNavigate for the /login destination.
func layoutImports(layout stml.LayoutSpec, emitLogout bool) []string {
	var imports []string
	if len(layout.NavItems) > 0 {
		imports = append(imports, "Link")
	}
	if layout.HasOutlet {
		imports = append(imports, "Outlet")
	}
	if emitLogout {
		imports = append(imports, "useNavigate")
	}
	return imports
}
