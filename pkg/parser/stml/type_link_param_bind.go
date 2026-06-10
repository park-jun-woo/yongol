//ff:type feature=stml-parse type=model
//ff:what LinkParamBind — data-link-params의 소스 → 대상 라우트 세그먼트 바인딩 한 건
package stml

// LinkParamBind is a single binding of a data-link-params attribute: the
// source value fills the named segment of the target page's resolved
// route (e.g. "item.id -> BuildingID"). Segment is empty for the elided
// form ("item.id"), legal only when the target route has exactly one
// required segment (TM-32 enforces the ambiguity rule).
type LinkParamBind struct {
	Source  string // "item.<Field>" (inside data-each) or "route.<Name>" (own page route)
	Segment string // target route segment name; empty = elided single-required-segment form
}
