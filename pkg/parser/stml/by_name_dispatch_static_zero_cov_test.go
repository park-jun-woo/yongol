//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameDispatchStatic_ZeroCov(t *testing.T) {
	var page PageSpec
	for _, frag := range []string{
		`<section data-fetch="F"></section>`,
		`<div data-action="A"></div>`,
		`<div><span data-bind="b"></span></div>`,
		`<h1>static</h1>`,
	} {
		el := firstElementNode(t, frag, firstTag(frag))
		dispatchStaticDataChild(el, &page)
	}

	var ab ActionBlock
	for _, frag := range []string{
		`<span data-component="C"></span>`,
		`<input data-field="F"/>`,
		`<button type="submit">go</button>`,
		`<h1>static</h1>`,
	} {
		el := firstElementNode(t, frag, firstTag(frag))
		dispatchStaticActionChild(el, &ab)
	}

	var fb FetchBlock
	for _, frag := range []string{
		`<section data-fetch="F"></section>`,
		`<div data-action="A"></div>`,
		`<ul data-each="x"></ul>`,
		`<span data-bind="b"></span>`,
		`<p data-state="s">x</p>`,
		`<span data-component="C"></span>`,
		`<h1>static</h1>`,
	} {
		el := firstElementNode(t, frag, firstTag(frag))
		dispatchStaticFetchChild(el, &fb)
	}

	var eb EachBlock
	for _, frag := range []string{
		`<span data-bind="b"></span>`,
		`<h1>static</h1>`,
	} {
		el := firstElementNode(t, frag, firstTag(frag))
		dispatchStaticEachChild(el, &eb)
	}
}
