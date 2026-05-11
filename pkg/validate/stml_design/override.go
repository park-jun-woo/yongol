//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what collectOverrides — STML 파일에서 <!-- @override --> 주석 직후 요소의 class를 수집
package stml_design

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// overrideSet maps filename → set of class attribute values that are overridden.
type overrideSet map[string]map[string]bool

// collectOverrides scans raw STML HTML files for <!-- @override --> comments and
// returns the set of class attribute values on elements immediately following such comments.
func collectOverrides(fs *yongol.Fullstack) overrideSet {
	result := make(overrideSet)
	frontendDir := filepath.Join(fs.SpecsDir, "frontend")

	for _, page := range fs.STMLPages {
		path := filepath.Join(frontendDir, page.FileName)
		classes := parseOverridesFromFile(path)
		if len(classes) > 0 {
			result[page.FileName] = classes
		}
	}
	return result
}

// parseOverridesFromFile reads an HTML file and finds class values on elements
// immediately preceded by <!-- @override --> comments.
func parseOverridesFromFile(path string) map[string]bool {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		return nil
	}

	classes := make(map[string]bool)
	walkForOverrides(doc, classes)
	return classes
}

// walkForOverrides recursively walks the DOM and finds elements preceded by
// an @override comment sibling.
func walkForOverrides(n *html.Node, classes map[string]bool) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.CommentNode && isOverrideComment(c.Data) {
			// Find next element sibling
			for next := c.NextSibling; next != nil; next = next.NextSibling {
				if next.Type == html.ElementNode {
					cls := getNodeAttr(next, "class")
					if cls != "" {
						classes[cls] = true
					}
					break
				}
				// Skip text nodes (whitespace)
				if next.Type == html.TextNode && strings.TrimSpace(next.Data) != "" {
					break
				}
			}
		}
		walkForOverrides(c, classes)
	}
}

// isOverrideComment checks if a comment node's data matches " @override ".
func isOverrideComment(data string) bool {
	return strings.TrimSpace(data) == "@override"
}

// getNodeAttr returns the value of the named attribute on an html.Node.
func getNodeAttr(n *html.Node, name string) string {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}

// isOverridden returns true if the given file+class combination is in the override set.
func isOverridden(ovr overrideSet, file, class string) bool {
	if class == "" {
		return false
	}
	m, ok := ovr[file]
	if !ok {
		return false
	}
	return m[class]
}
