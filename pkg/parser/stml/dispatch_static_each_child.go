//ff:func feature=stml-parse type=parser control=selection
//ff:what each 항목 내 정적 래퍼의 자식을 분기 처리
package stml

import "golang.org/x/net/html"

func dispatchStaticEachChild(c *html.Node, eb *EachBlock) ChildNode {
	switch {
	case getAttr(c, "data-bind") != "":
		bf := getAttr(c, "data-bind")
		bind := FieldBind{Name: bf, Tag: c.Data, Type: getAttr(c, "type"), ClassName: getAttr(c, "class")}
		eb.Binds = append(eb.Binds, bind)
		return ChildNode{Kind: "bind", Bind: &bind}
	default:
		childSe := parseStaticElement(c)
		return ChildNode{Kind: "static", Static: &childSe}
	}
}
