//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what collectOverrides — STML 파일에서 <!-- @override class="..." --> 주석의 class 값을 수집
package stml_design

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// overrideSet maps filename → set of class attribute values that are overridden.
type overrideSet map[string]map[string]bool

// overrideClassRe extracts the class value from @override comment text.
// Matches: @override class="..." or @override class='...'
var overrideClassRe = regexp.MustCompile(`@override\s+class=["']([^"']+)["']`)

// collectOverrides scans raw STML HTML files for <!-- @override class="..." --> comments
// and returns the set of class attribute values extracted from those comments.
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

// parseOverridesFromFile reads an HTML file and extracts class values from
// <!-- @override class="..." --> comments.
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

// walkForOverrides recursively walks the DOM and extracts class values from
// @override comment nodes.
func walkForOverrides(n *html.Node, classes map[string]bool) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.CommentNode && isOverrideComment(c.Data) {
			cls := extractOverrideClass(c.Data)
			if cls != "" {
				classes[cls] = true
			}
		}
		walkForOverrides(c, classes)
	}
}

// isOverrideComment checks if a comment node's data starts with "@override".
// Matches both <!-- @override --> and <!-- @override class="..." -->.
func isOverrideComment(data string) bool {
	return strings.HasPrefix(strings.TrimSpace(data), "@override")
}

// extractOverrideClass extracts the class value from an @override comment.
// Returns "" if no class attribute is present (structure-only override).
func extractOverrideClass(data string) string {
	m := overrideClassRe.FindStringSubmatch(data)
	if m == nil {
		return ""
	}
	return m[1]
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
