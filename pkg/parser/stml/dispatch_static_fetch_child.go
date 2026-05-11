//ff:func feature=stml-parse type=parser control=selection
//ff:what 정적 fetch 래퍼 내 단일 자식 요소를 분기 처리
package stml

import "golang.org/x/net/html"

// dispatchStaticFetchChild handles a single child element inside a static fetch wrapper.
func dispatchStaticFetchChild(c *html.Node, fb *FetchBlock) *ChildNode {
	switch {
	case getAttr(c, "data-fetch") != "":
		child := parseFetchBlock(c, getAttr(c, "data-fetch"))
		fb.NestedFetches = append(fb.NestedFetches, child)
		return &ChildNode{Kind: "fetch", Fetch: &child}
	case getAttr(c, "data-action") != "":
		ab := parseActionBlock(c, getAttr(c, "data-action"))
		return &ChildNode{Kind: "action", Action: &ab}
	case getAttr(c, "data-each") != "":
		eb := parseEachBlock(c, getAttr(c, "data-each"))
		fb.Eaches = append(fb.Eaches, eb)
		return &ChildNode{Kind: "each", Each: &eb}
	case getAttr(c, "data-bind") != "":
		field := getAttr(c, "data-bind")
		bind := FieldBind{Name: field, Tag: c.Data, Type: getAttr(c, "type"), ClassName: getAttr(c, "class")}
		fb.Binds = append(fb.Binds, bind)
		return &ChildNode{Kind: "bind", Bind: &bind}
	case getAttr(c, "data-state") != "":
		sb := parseStateBind(c, getAttr(c, "data-state"))
		fb.States = append(fb.States, sb)
		return &ChildNode{Kind: "state", State: &sb}
	case getAttr(c, "data-component") != "":
		comp := getAttr(c, "data-component")
		cr := ComponentRef{Name: comp, Bind: getAttr(c, "data-bind"), Field: getAttr(c, "data-field"), ClassName: getAttr(c, "class")}
		fb.Components = append(fb.Components, cr)
		return &ChildNode{Kind: "component", Component: &cr}
	case hasContent(c) || hasDescendantDataInFetch(c):
		childStatic := parseStaticInFetch(c, fb)
		return &ChildNode{Kind: "static", Static: &childStatic}
	default:
		return nil
	}
}
