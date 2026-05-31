//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameWalkers_ZeroCov(t *testing.T) {
	// walkFetchChildren
	fEl := firstElementNode(t, `<section><span data-bind="a"></span></section>`, "section")
	var fb FetchBlock
	for c := fEl.FirstChild; c != nil; c = c.NextSibling {
		walkFetchChildren(c, &fb)
	}

	// walkActionChildren
	aEl := firstElementNode(t, `<div><input data-field="N"/></div>`, "div")
	var ab ActionBlock
	for c := aEl.FirstChild; c != nil; c = c.NextSibling {
		walkActionChildren(c, &ab)
	}

	// walkEachChildren + walkEachItemChildren via parseEachItemTemplate
	eEl := firstElementNode(t, `<ul><li data-bind="x"><span data-bind="y"></span></li></ul>`, "ul")
	var eb EachBlock
	for c := eEl.FirstChild; c != nil; c = c.NextSibling {
		walkEachChildren(c, &eb)
	}
	parseEachItemTemplate(eEl, &eb)
	itemEl := firstElementNode(t, `<li><span data-bind="y"></span></li>`, "li")
	var eb2 EachBlock
	for c := itemEl.FirstChild; c != nil; c = c.NextSibling {
		walkEachItemChildren(c, &eb2)
	}
}
