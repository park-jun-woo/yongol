//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what data-state 요소에서 StateBind 구성
package stml

import "golang.org/x/net/html"

// parseStateBind builds a StateBind from a data-state element.
func parseStateBind(n *html.Node, condition string) StateBind {
	sb := StateBind{
		Tag:       n.Data,
		ClassName: getAttr(n, "class"),
		Condition: condition,
		Text:      directText(n),
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		if op := getAttr(c, "data-action"); op != "" {
			ab := parseActionBlock(c, op)
			sb.Children = append(sb.Children, ChildNode{Kind: "action", Action: &ab})
		} else if hasContent(c) {
			se := parseStaticElement(c)
			sb.Children = append(sb.Children, ChildNode{Kind: "static", Static: &se})
		}
	}
	return sb
}
