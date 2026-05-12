//ff:type feature=validate type=model
//ff:what tokenRef — STML class 속성에서 추출된 디자인 토큰 참조 위치 정보
package stml_design

// tokenRef records where a potential design token name appears in STML.
type tokenRef struct {
	File  string // STML filename
	Class string // full class attribute string
	Name  string // extracted token name (e.g. "primary", "sm")
}
