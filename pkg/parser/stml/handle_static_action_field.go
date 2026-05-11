//ff:func feature=stml-parse type=parser control=sequence
//ff:what 정적 action 래퍼 내 data-field 요소 처리
package stml

import "golang.org/x/net/html"

func handleStaticActionField(c *html.Node, ab *ActionBlock) *ChildNode {
	field := getAttr(c, "data-field")
	bind := FieldBind{Name: field, Tag: c.Data, Type: getAttr(c, "type"), ClassName: getAttr(c, "class"), Placeholder: getAttr(c, "placeholder")}
	ab.Fields = append(ab.Fields, bind)
	return &ChildNode{Kind: "bind", Bind: &bind}
}
