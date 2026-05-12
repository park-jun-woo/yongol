//ff:type feature=stml-parse type=model
//ff:what NavItem — 레이아웃 data-nav 속성에서 추출한 네비게이션 링크 구조체

package stml

// NavItem represents a navigation link extracted from a data-nav attribute.
type NavItem struct {
	Path  string // data-nav attribute value (e.g., "/workflows")
	Label string // link text content
}
