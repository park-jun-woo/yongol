//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what parseOverridesFromFile — HTML 파일에서 @override class 주석 값 추출
package stml_design

import (
	"os"

	"golang.org/x/net/html"
)

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
