//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what fetch 블록 내 data-component 요소 처리
package stml

import "golang.org/x/net/html"

func handleFetchComponent(n *html.Node, fb *FetchBlock) bool {
	comp := getAttr(n, "data-component")
	cr := ComponentRef{
		Name:      comp,
		Bind:      getAttr(n, "data-bind"),
		Field:     getAttr(n, "data-field"),
		ClassName: getAttr(n, "class"),
	}
	fb.Components = append(fb.Components, cr)
	fb.Children = append(fb.Children, ChildNode{Kind: "component", Component: &cr})
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkFetchChildren(c, fb)
	}
	return true
}
