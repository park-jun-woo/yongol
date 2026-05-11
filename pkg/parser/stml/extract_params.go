//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 요소에서 data-param-* 속성을 추출
package stml

import (
	"strings"

	"golang.org/x/net/html"
)

// extractParams extracts data-param-* attributes from an element.
func extractParams(n *html.Node) []ParamBind {
	var params []ParamBind
	for _, attr := range n.Attr {
		if strings.HasPrefix(attr.Key, "data-param-") {
			paramName := attr.Key[len("data-param-"):]
			params = append(params, ParamBind{
				Name:   kebabToCamel(paramName),
				Source: attr.Val,
			})
		}
	}
	return params
}
