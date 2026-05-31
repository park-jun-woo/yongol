//ff:func feature=stml-parse type=test control=sequence
//ff:what isImplicitTag/isDataElement/getAttr/hasAttr/hasDataAttr/hasFieldAttr/splitTrim/extractAllText/collectText/directText/hasContent/extractNonEmptyText
package stml

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func firstElementNode(t *testing.T, fragment, tag string) *html.Node {
	t.Helper()
	doc, err := html.Parse(strings.NewReader(fragment))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	var found *html.Node
	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if found != nil {
			return
		}
		if n.Type == html.ElementNode && n.Data == tag {
			found = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)
	if found == nil {
		t.Fatalf("tag %q not found in %q", tag, fragment)
	}
	return found
}
