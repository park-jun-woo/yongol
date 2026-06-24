//ff:func feature=stml-parse type=parser control=sequence dimension=1
//ff:what data-component 요소에서 ComponentRef 구성
package stml

import "golang.org/x/net/html"

// buildComponentRefFromEach builds a ComponentRef from a data-component element.
func buildComponentRefFromEach(n *html.Node) ComponentRef {
	return ComponentRef{
		Name:      getAttr(n, "data-component"),
		Bind:      getAttr(n, "data-bind"),
		Field:     getAttr(n, "data-field"),
		ClassName: getAttr(n, "class"),
		Variant:   getAttr(n, "data-variant"),
		Size:      getAttr(n, "data-size"),
	}
}
