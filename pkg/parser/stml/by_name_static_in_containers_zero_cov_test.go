//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameStaticInContainers_ZeroCov(t *testing.T) {
	actEl := firstElementNode(t,
		`<div class="wrap">label<input data-field="F"/><span data-component="C"></span><span>txt</span></div>`, "div")
	var ab ActionBlock
	seA := parseStaticInAction(actEl, &ab)
	if seA.Tag != "div" {
		t.Errorf("parseStaticInAction Tag = %q", seA.Tag)
	}

	eachEl := firstElementNode(t, `<div class="w">x<span data-bind="b"></span><span>txt</span></div>`, "div")
	var eb EachBlock
	seE := parseStaticInEach(eachEl, &eb)
	if seE.Tag != "div" {
		t.Errorf("parseStaticInEach Tag = %q", seE.Tag)
	}

	fetchEl := firstElementNode(t,
		`<div class="w">x<span data-bind="b"></span><section data-fetch="N"></section><span>txt</span></div>`, "div")
	var fb FetchBlock
	seF := parseStaticInFetch(fetchEl, &fb)
	if seF.Tag != "div" {
		t.Errorf("parseStaticInFetch Tag = %q", seF.Tag)
	}

	dataEl := firstElementNode(t,
		`<div><section data-fetch="F"></section><div data-action="A"></div><div><span data-bind="b"></span></div><h1>x</h1></div>`, "div")
	var page PageSpec
	seD := parseStaticWithDataChildren(dataEl, &page)
	if seD.Tag != "div" {
		t.Errorf("parseStaticWithDataChildren Tag = %q", seD.Tag)
	}

	// walkStaticActionChild + walkStaticActionNestedChildren
	wrapEl := firstElementNode(t,
		`<div><input data-field="F"/><span data-component="C"></span><div><input data-field="G"/></div></div>`, "div")
	var ab3 ActionBlock
	se3 := StaticElement{Tag: "div"}
	for c := wrapEl.FirstChild; c != nil; c = c.NextSibling {
		walkStaticActionChild(c, &ab3, &se3)
	}
	nestEl := firstElementNode(t, `<div><input data-field="F"/></div>`, "div")
	var ab4 ActionBlock
	walkStaticActionNestedChildren(nestEl, &ab4)
}
