//ff:func feature=stml-parse type=parser control=selection
//ff:what each 블록 내 단일 요소를 분기 처리
package stml

import "golang.org/x/net/html"

func dispatchEachChild(n *html.Node, eb *EachBlock) bool {
	switch {
	case getAttr(n, "data-bind") != "":
		field := getAttr(n, "data-bind")
		bind := FieldBind{
			Name:      field,
			Tag:       n.Data,
			Type:      getAttr(n, "type"),
			ClassName: getAttr(n, "class"),
		}
		eb.Binds = append(eb.Binds, bind)
		eb.Children = append(eb.Children, ChildNode{Kind: "bind", Bind: &bind})
		return true
	case getAttr(n, "data-state") != "":
		cond := getAttr(n, "data-state")
		sb := parseStateBind(n, cond)
		eb.States = append(eb.States, sb)
		eb.Children = append(eb.Children, ChildNode{Kind: "state", State: &sb})
		return true
	case getAttr(n, "data-component") != "":
		comp := getAttr(n, "data-component")
		cr := ComponentRef{
			Name:      comp,
			Bind:      getAttr(n, "data-bind"),
			Field:     getAttr(n, "data-field"),
			ClassName: getAttr(n, "class"),
		}
		eb.Components = append(eb.Components, cr)
		eb.Children = append(eb.Children, ChildNode{Kind: "component", Component: &cr})
		return true
	case hasContent(n):
		se := parseStaticInEach(n, eb)
		eb.Children = append(eb.Children, ChildNode{Kind: "static", Static: &se})
		return true
	default:
		return false
	}
}
