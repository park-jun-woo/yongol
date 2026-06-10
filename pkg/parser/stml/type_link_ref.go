//ff:type feature=stml-parse type=model
//ff:what LinkRef — data-link 요소(대상 페이지명 참조 + 세그먼트 파라미터 매핑)를 나타내는 구조체
package stml

// LinkRef represents a data-link element: a declaration that clicking this
// element navigates to another STML page. TargetPage is a page-name
// reference (STML filename without .html), not a path literal — route
// paths are a derived projection (RoutePaths), so the SSOT records only
// the decision of which page. Params bind the target route's segments
// from the row context (item.<Field>) or the current route (route.<Name>).
type LinkRef struct {
	Tag        string          // original HTML tag (e.g. "a", "li")
	ClassName  string          // class attribute value
	Text       string          // direct text content
	TargetPage string          // data-link value: target STML page name (filename without .html)
	ParamsRaw  string          // raw data-link-params value (TM-32 re-parses it for syntax diagnostics)
	Params     []LinkParamBind // parsed bindings (empty when absent or syntactically invalid)
	Children   []ChildNode     // children in DOM order (bind/static render reuse)

	// TargetPattern is codegen-populated (pkg/generate/react/stml, like
	// EachBlock.KeyField): the target page's resolved route pattern
	// (stml.RoutePaths first pattern). The parser always leaves it empty.
	TargetPattern string
}
