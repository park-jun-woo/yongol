//ff:type feature=stml-parse type=model
//ff:what 페이지가 소비하는 route.<Name> 파라미터 (이름 + 필수 여부)
package stml

// routeParam is a route.<Name> parameter a page consumes. Required params
// are consumed by some data-fetch (the page cannot render without them);
// optional params are consumed only by data-action blocks (the page must
// stay reachable without them).
type routeParam struct {
	Name     string // segment name as declared (e.g. "BuildingID")
	Required bool   // true = fetch-consumed (":Name"), false = action-only (":Name?")
}
