//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what scanInlineStyles — HTML 파일 파싱 후 inline style 속성의 하드코딩 색상 검사
package stml_design

import (
	"os"

	"golang.org/x/net/html"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanInlineStyles parses an HTML file and checks style attributes for hardcoded colors.
func scanInlineStyles(path, filename string, hexToToken map[string]string, ovr overrideSet) []diagnostic.Diagnostic {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	walkInlineStyles(doc, filename, hexToToken, ovr, &diags)
	return diags
}
