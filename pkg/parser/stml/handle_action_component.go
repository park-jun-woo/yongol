//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what action 블록 내 data-component 요소 처리
package stml

import "golang.org/x/net/html"

func handleActionComponent(n *html.Node, ab *ActionBlock) bool {
	comp := getAttr(n, "data-component")
	if f := getAttr(n, "data-field"); f != "" {
		bind := FieldBind{
			Name:      f,
			Tag:       "data-component:" + comp,
			ClassName: getAttr(n, "class"),
		}
		ab.Fields = append(ab.Fields, bind)
		ab.Children = append(ab.Children, ChildNode{Kind: "bind", Bind: &bind})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkActionChildren(c, ab)
	}
	return true
}
