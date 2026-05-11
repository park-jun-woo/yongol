//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 정적 action 래퍼 내 data-component 요소를 재귀 처리
package stml

import "golang.org/x/net/html"

func handleWalkStaticActionComponent(c *html.Node, ab *ActionBlock, se *StaticElement) {
	comp := getAttr(c, "data-component")
	if f := getAttr(c, "data-field"); f != "" {
		bind := FieldBind{Name: f, Tag: "data-component:" + comp, ClassName: getAttr(c, "class")}
		ab.Fields = append(ab.Fields, bind)
		if se != nil {
			se.Children = append(se.Children, ChildNode{Kind: "bind", Bind: &bind})
		}
	}
	for gc := c.FirstChild; gc != nil; gc = gc.NextSibling {
		if gc.Type == html.ElementNode {
			walkStaticActionChild(gc, ab, se)
		}
	}
}
