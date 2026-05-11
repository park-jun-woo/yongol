//ff:func feature=stml-parse type=util control=sequence
//ff:what 요소에 data-field 또는 data-component+data-field 속성이 있는지 확인
package stml

import "golang.org/x/net/html"

func hasFieldAttr(c *html.Node) bool {
	if getAttr(c, "data-field") != "" {
		return true
	}
	return getAttr(c, "data-component") != "" && getAttr(c, "data-field") != ""
}
