//ff:type feature=stml-parse type=model
//ff:what data-param-* 속성을 나타내는 구조체
package stml

// ParamBind represents a data-param-* attribute.
type ParamBind struct {
	Name   string // parameter name (e.g. "ReservationID")
	Source string // value source (e.g. "route.ReservationID")
	// Optional marks a route.<Name> source whose URL segment is optional
	// (":Name?") — consumed only by data-action blocks, never by a data-fetch
	// (collectRouteParams's Required=false). Optional integer params need a
	// null guard before Number() so an absent segment does not send NaN
	// (BUG-136). Populated by the react emitter (markOptionalRouteParams).
	Optional bool
}
