//ff:func feature=stml-parse type=parser control=sequence
//ff:what action 블록 내 data-field 요소 처리
package stml

import "golang.org/x/net/html"

func handleActionField(n *html.Node, ab *ActionBlock) bool {
	field := getAttr(n, "data-field")
	bind := FieldBind{
		Name:        field,
		Tag:         n.Data,
		Type:        getAttr(n, "type"),
		ClassName:   getAttr(n, "class"),
		Placeholder: getAttr(n, "placeholder"),
	}
	ab.Fields = append(ab.Fields, bind)
	ab.Children = append(ab.Children, ChildNode{Kind: "bind", Bind: &bind})
	return true
}
