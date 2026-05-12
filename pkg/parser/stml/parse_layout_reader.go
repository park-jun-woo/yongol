//ff:func feature=stml-parse type=parser control=sequence
//ff:what ParseLayoutReader — io.Reader 에서 레이아웃 HTML 파싱하여 LayoutSpec 반환

package stml

import (
	"io"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"golang.org/x/net/html"
)

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
