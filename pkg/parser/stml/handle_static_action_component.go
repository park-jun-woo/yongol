//ff:func feature=stml-parse type=parser control=sequence
//ff:what 정적 action 래퍼 내 data-component 요소 처리
package stml

import "golang.org/x/net/html"

func handleStaticActionComponent(c *html.Node, ab *ActionBlock) *ChildNode {
	comp := getAttr(c, "data-component")
	if f := getAttr(c, "data-field"); f != "" {
		bind := FieldBind{Name: f, Tag: "data-component:" + comp, ClassName: getAttr(c, "class")}
		ab.Fields = append(ab.Fields, bind)
		walkStaticActionNestedChildren(c, ab)
		return &ChildNode{Kind: "bind", Bind: &bind}
	}
	walkStaticActionNestedChildren(c, ab)
	return nil
}
