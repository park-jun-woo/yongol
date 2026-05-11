//ff:func feature=stml-parse type=parser control=selection
//ff:what action 블록 내 단일 요소를 분기 처리
package stml

import "golang.org/x/net/html"

// dispatchActionChild handles a single element node inside an action block.
func dispatchActionChild(n *html.Node, ab *ActionBlock) bool {
	switch {
	case getAttr(n, "data-component") != "":
		return handleActionComponent(n, ab)
	case getAttr(n, "data-field") != "":
		return handleActionField(n, ab)
	case n.Data == "button" && getAttr(n, "type") == "submit":
		ab.SubmitText = directText(n)
		return true
	case hasContent(n) || hasDescendantField(n):
		se := parseStaticInAction(n, ab)
		ab.Children = append(ab.Children, ChildNode{Kind: "static", Static: &se})
		return true
	default:
		return false
	}
}
