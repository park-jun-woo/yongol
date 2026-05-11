//ff:func feature=stml-gen type=util control=sequence
//ff:what 페이지의 루트 엘리먼트 태그와 클래스명을 결정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func findRootElement(page stmlparser.PageSpec) (string, string) {
	if len(page.Children) == 1 && page.Children[0].Kind == "static" {
		se := page.Children[0].Static
		return se.Tag, se.ClassName
	}
	return "div", ""
}
