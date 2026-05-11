//ff:func feature=stml-parse type=parser control=sequence
//ff:what io.Reader에서 HTML을 파싱하여 PageSpec 반환
package stml

import (
	"io"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"golang.org/x/net/html"
)

// ParseReader parses HTML from a reader and returns a PageSpec.
func ParseReader(filename string, r io.Reader) (PageSpec, []diagnostic.Diagnostic) {
	doc, err := html.Parse(r)
	if err != nil {
		return PageSpec{}, []diagnostic.Diagnostic{{
			File:    filename,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "html parse: " + err.Error(),
		}}
	}

	name := strings.TrimSuffix(filename, ".html")
	page := PageSpec{
		Name:     name,
		FileName: filename,
	}

	walkTopLevel(doc, &page)
	return page, nil
}
