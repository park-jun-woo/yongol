//ff:type feature=stml-parse type=model
//ff:what STML 페이지 파싱 결과를 나타내는 구조체
package stml

// PageSpec represents a single STML page parsed from an HTML file.
type PageSpec struct {
	Name     string        // page name derived from filename (e.g. "login-page")
	FileName string        // original filename (e.g. "login-page.html")
	Fetches  []FetchBlock  // data-fetch blocks (for validation)
	Actions  []ActionBlock // data-action blocks (for validation)
	Children []ChildNode   // all top-level children in DOM order (for codegen)
}
