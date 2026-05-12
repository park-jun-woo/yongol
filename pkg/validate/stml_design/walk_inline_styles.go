//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what walkInlineStyles — DOM 순회하며 inline style 속성의 하드코딩 색상 검출
package stml_design

import (
	"golang.org/x/net/html"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// walkInlineStyles recursively walks the DOM checking for inline style hardcoded colors.
func walkInlineStyles(n *html.Node, filename string, hexToToken map[string]string, ovr overrideSet, diags *[]diagnostic.Diagnostic) {
	if n.Type == html.ElementNode {
		style := getNodeAttr(n, "style")
		if style != "" && !isPrecededByOverride(n) {
			checkStyleColors(style, filename, hexToToken, diags)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkInlineStyles(c, filename, hexToToken, ovr, diags)
	}
}
