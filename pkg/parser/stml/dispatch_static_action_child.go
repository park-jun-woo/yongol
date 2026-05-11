//ff:func feature=stml-parse type=parser control=selection
//ff:what 정적 action 래퍼 내 단일 자식 요소를 분기 처리
package stml

import "golang.org/x/net/html"

// dispatchStaticActionChild handles a single child element inside a static action wrapper.
func dispatchStaticActionChild(c *html.Node, ab *ActionBlock) *ChildNode {
	switch {
	case getAttr(c, "data-component") != "":
		return handleStaticActionComponent(c, ab)
	case getAttr(c, "data-field") != "":
		return handleStaticActionField(c, ab)
	case c.Data == "button" && getAttr(c, "type") == "submit":
		ab.SubmitText = directText(c)
		return nil
	case hasContent(c) || hasDescendantField(c):
		childSe := parseStaticInAction(c, ab)
		return &ChildNode{Kind: "static", Static: &childSe}
	default:
		return nil
	}
}
