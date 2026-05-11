//ff:func feature=stml-parse type=parser control=sequence
//ff:what 레이아웃 HTML 파일을 파싱하여 LayoutSpec 반환 (data-nav, slot data-outlet 추출)
package stml

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"golang.org/x/net/html"
)

// ParseLayoutFile parses a single layout HTML file and returns a LayoutSpec.
func ParseLayoutFile(path string) (LayoutSpec, []diagnostic.Diagnostic) {
	f, err := os.Open(path)
	if err != nil {
		return LayoutSpec{}, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: err.Error(),
		}}
	}
	defer f.Close()

	return ParseLayoutReader(filepath.Base(path), path, f)
}

// ParseLayoutReader parses layout HTML from a reader and returns a LayoutSpec.
func ParseLayoutReader(filename, filePath string, r io.Reader) (LayoutSpec, []diagnostic.Diagnostic) {
	doc, err := html.Parse(r)
	if err != nil {
		return LayoutSpec{}, []diagnostic.Diagnostic{{
			File:    filename,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "html parse: " + err.Error(),
		}}
	}

	name := strings.TrimSuffix(filename, ".html")
	layout := LayoutSpec{
		Name: name,
		File: filePath,
	}

	walkLayoutNode(doc, &layout)
	return layout, nil
}

// ParseLayoutDir parses all .html files in the given layouts directory
// and returns a LayoutSpec for each.
func ParseLayoutDir(dir string) ([]LayoutSpec, []diagnostic.Diagnostic) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("read layout dir %s: %s", dir, err),
		}}
	}

	var layouts []LayoutSpec
	var allDiags []diagnostic.Diagnostic
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".html") {
			continue
		}
		layout, diags := ParseLayoutFile(filepath.Join(dir, e.Name()))
		if len(diags) > 0 {
			allDiags = append(allDiags, diags...)
			continue
		}
		layouts = append(layouts, layout)
	}
	if len(allDiags) > 0 {
		return nil, allDiags
	}
	return layouts, nil
}

// walkLayoutNode recursively walks the DOM tree to extract data-nav links
// and slot data-outlet elements.
func walkLayoutNode(n *html.Node, layout *LayoutSpec) {
	if n.Type == html.ElementNode {
		// Check for data-nav attribute on <a> elements.
		if n.Data == "a" {
			if nav := getAttr(n, "data-nav"); nav != "" {
				label := extractAllText(n)
				layout.NavItems = append(layout.NavItems, NavItem{
					Path:  nav,
					Label: label,
				})
			}
		}
		// Check for <slot data-outlet />.
		if n.Data == "slot" && hasAttr(n, "data-outlet") {
			layout.HasOutlet = true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkLayoutNode(c, layout)
	}
}

// extractAllText collects all text content from an element and its descendants.
func extractAllText(n *html.Node) string {
	var sb strings.Builder
	collectText(n, &sb)
	return strings.TrimSpace(sb.String())
}

// collectText recursively collects text node content.
func collectText(n *html.Node, sb *strings.Builder) {
	if n.Type == html.TextNode {
		sb.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, sb)
	}
}
