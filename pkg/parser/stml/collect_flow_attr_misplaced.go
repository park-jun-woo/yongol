//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what DOM 전체를 순회하며 허용 위치 밖의 흐름 속성을 PageSpec에 기록 (TM-25 근거)
package stml

import "golang.org/x/net/html"

// collectFlowAttrMisplaced walks the whole DOM recording flow attributes on
// illegal positions into page.FlowAttrMisplaced: data-capture and
// data-redirect are valid only on a data-action element itself, and
// data-on-error is valid only on an element inside a data-action block
// (the action root does not count as inside).
func collectFlowAttrMisplaced(n *html.Node, inAction bool, page *PageSpec) {
	if n.Type == html.ElementNode {
		isAction := getAttr(n, "data-action") != ""
		if !isAction && hasAttr(n, "data-capture") {
			page.FlowAttrMisplaced = append(page.FlowAttrMisplaced, FlowAttrMisplaced{Attr: "data-capture", Tag: n.Data})
		}
		if !isAction && hasAttr(n, "data-redirect") {
			page.FlowAttrMisplaced = append(page.FlowAttrMisplaced, FlowAttrMisplaced{Attr: "data-redirect", Tag: n.Data})
		}
		if !inAction && hasAttr(n, "data-on-error") {
			page.FlowAttrMisplaced = append(page.FlowAttrMisplaced, FlowAttrMisplaced{Attr: "data-on-error", Tag: n.Data})
		}
		inAction = inAction || isAction
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectFlowAttrMisplaced(c, inAction, page)
	}
}
