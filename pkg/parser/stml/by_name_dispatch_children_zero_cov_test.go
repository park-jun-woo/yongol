//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameDispatchChildren_ZeroCov(t *testing.T) {
	// dispatchFetchChild across branches
	var fb FetchBlock
	for _, frag := range []string{
		`<section data-fetch="Sub"></section>`,
		`<div data-action="A"></div>`,
		`<ul data-each="items"></ul>`,
		`<span data-bind="b"></span>`,
		`<p data-state="s">x</p>`,
		`<span data-component="C"></span>`,
		`<h1>static</h1>`,
	} {
		el := firstElementNode(t, frag, firstTag(frag))
		dispatchFetchChild(el, &fb)
	}

	// dispatchActionChild
	var ab ActionBlock
	for _, frag := range []string{
		`<span data-component="C"></span>`,
		`<input data-field="F"/>`,
		`<button type="submit">go</button>`,
		`<h1>static</h1>`,
	} {
		el := firstElementNode(t, frag, firstTag(frag))
		dispatchActionChild(el, &ab)
	}

	// dispatchEachChild
	var eb EachBlock
	for _, frag := range []string{
		`<button data-action="A">x</button>`,
		`<span data-bind="b"></span>`,
		`<p data-state="s">x</p>`,
		`<span data-component="C"></span>`,
		`<h1>static</h1>`,
	} {
		el := firstElementNode(t, frag, firstTag(frag))
		dispatchEachChild(el, &eb)
	}
}
