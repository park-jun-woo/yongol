//ff:func feature=stml-parse type=parser control=selection
//ff:what 정적 action 래퍼 내 단일 자식 요소를 처리
package stml

import "golang.org/x/net/html"

// walkStaticActionChild handles a single child element inside a static action wrapper.
func walkStaticActionChild(c *html.Node, ab *ActionBlock, se *StaticElement) {
	switch {
	case getAttr(c, "data-component") != "":
		handleWalkStaticActionComponent(c, ab, se)
	case getAttr(c, "data-field") != "":
		handleWalkStaticActionField(c, ab, se)
	case c.Data == "button" && getAttr(c, "type") == "submit":
		ab.SubmitText = directText(c)
	}
}
