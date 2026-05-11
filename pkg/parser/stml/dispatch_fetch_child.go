//ff:func feature=stml-parse type=parser control=selection
//ff:what fetch 블록 내 단일 요소를 분기 처리
package stml

import "golang.org/x/net/html"

// dispatchFetchChild handles a single element node inside a fetch block.
func dispatchFetchChild(n *html.Node, fb *FetchBlock) bool {
	switch {
	case getAttr(n, "data-fetch") != "":
		child := parseFetchBlock(n, getAttr(n, "data-fetch"))
		fb.NestedFetches = append(fb.NestedFetches, child)
		fb.Children = append(fb.Children, ChildNode{Kind: "fetch", Fetch: &child})
		return true
	case getAttr(n, "data-action") != "":
		ab := parseActionBlock(n, getAttr(n, "data-action"))
		fb.Children = append(fb.Children, ChildNode{Kind: "action", Action: &ab})
		return true
	case getAttr(n, "data-each") != "":
		eb := parseEachBlock(n, getAttr(n, "data-each"))
		fb.Eaches = append(fb.Eaches, eb)
		fb.Children = append(fb.Children, ChildNode{Kind: "each", Each: &eb})
		return true
	case getAttr(n, "data-bind") != "":
		field := getAttr(n, "data-bind")
		bind := FieldBind{
			Name:      field,
			Tag:       n.Data,
			Type:      getAttr(n, "type"),
			ClassName: getAttr(n, "class"),
		}
		fb.Binds = append(fb.Binds, bind)
		fb.Children = append(fb.Children, ChildNode{Kind: "bind", Bind: &bind})
		return true
	case getAttr(n, "data-state") != "":
		sb := parseStateBind(n, getAttr(n, "data-state"))
		fb.States = append(fb.States, sb)
		fb.Children = append(fb.Children, ChildNode{Kind: "state", State: &sb})
		return true
	case getAttr(n, "data-component") != "":
		return handleFetchComponent(n, fb)
	case hasContent(n) || hasDescendantDataInFetch(n):
		se := parseStaticInFetch(n, fb)
		fb.Children = append(fb.Children, ChildNode{Kind: "static", Static: &se})
		return true
	default:
		return false
	}
}
