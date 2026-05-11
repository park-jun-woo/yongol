//ff:type feature=stml-parse type=model
//ff:what data-param-* 속성을 나타내는 구조체
package stml

// ParamBind represents a data-param-* attribute.
type ParamBind struct {
	Name   string // parameter name (e.g. "ReservationID")
	Source string // value source (e.g. "route.ReservationID")
}
