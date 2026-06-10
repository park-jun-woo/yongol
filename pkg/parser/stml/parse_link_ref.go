//ff:func feature=stml-parse type=parser control=sequence
//ff:what data-link 요소의 속성(대상 페이지·파라미터 매핑)에서 LinkRef 구성 (자식은 컨텍스트별 호출자가 파싱)
package stml

import "golang.org/x/net/html"

// parseLinkRef builds a LinkRef from a data-link element's attributes.
// Children are parsed by the context-specific callers (each / fetch /
// static) so descendant binds register with the enclosing block for
// validation.
func parseLinkRef(n *html.Node) LinkRef {
	lr := LinkRef{
		Tag:        n.Data,
		ClassName:  getAttr(n, "class"),
		Text:       directText(n),
		TargetPage: getAttr(n, "data-link"),
		ParamsRaw:  getAttr(n, "data-link-params"),
	}
	if lr.ParamsRaw != "" {
		// Syntax errors are surfaced by TM-32 at validate time (re-parse).
		lr.Params, _ = ParseLinkParams(lr.ParamsRaw)
	}
	return lr
}
