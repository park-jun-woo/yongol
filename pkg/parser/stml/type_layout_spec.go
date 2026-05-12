//ff:type feature=stml-parse type=model
//ff:what 레이아웃 STML 파싱 결과를 나타내는 구조체
package stml

// LayoutSpec represents a single STML layout parsed from an HTML file
// in the specs/frontend/layouts/ directory.
type LayoutSpec struct {
	Name      string    // layout name derived from filename (e.g., "app")
	File      string    // original file path
	NavItems  []NavItem // data-nav navigation links
	HasOutlet bool      // whether slot data-outlet exists
}
