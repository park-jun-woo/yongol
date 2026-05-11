//ff:func feature=stml-parse type=util control=sequence
//ff:what 요소가 data-fetch 또는 data-action 속성을 가진 ElementNode인지 판별
package stml

import "golang.org/x/net/html"

func isDataElement(c *html.Node) bool {
	return c.Type == html.ElementNode &&
		(getAttr(c, "data-fetch") != "" || getAttr(c, "data-action") != "")
}
