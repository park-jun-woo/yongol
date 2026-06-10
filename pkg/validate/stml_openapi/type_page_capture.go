//ff:type feature=validate type=model topic=stml-openapi
//ff:what pageCapture — data-capture 바인딩과 그것이 선언된 페이지 파일명 쌍

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// pageCapture pairs a parsed data-capture binding with the STML page file
// that declares it, so project-level rules (TM-21/22/24) can point their
// diagnostics at the declaring page.
type pageCapture struct {
	File string
	Bind stml.CaptureBind
}
