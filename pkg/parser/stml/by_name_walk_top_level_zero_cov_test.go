//ff:func feature=stml-parse type=test control=sequence
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

func TestByNameWalkTopLevel_ZeroCov(t *testing.T) {
	frag := `<main>
  <section data-fetch="ListItems"><span data-bind="title"></span></section>
  <div data-action="Create"><input data-field="Name" /><button type="submit">Go</button></div>
  <div><span data-bind="nested"></span></div>
  <h1>Hello</h1>
</main>`
	n := firstElementNode(t, frag, "main")
	var page PageSpec
	walkTopLevel(n, &page)
	if len(page.Children) == 0 {
		t.Fatalf("walkTopLevel produced no children")
	}

	// dispatchTopLevelElement direct on a fetch element.
	fetchEl := firstElementNode(t, `<section data-fetch="X"><span data-bind="a"></span></section>`, "section")
	var p2 PageSpec
	if !dispatchTopLevelElement(fetchEl, &p2) {
		t.Errorf("dispatchTopLevelElement(fetch) = false")
	}
	// action element.
	actEl := firstElementNode(t, `<div data-action="A"><input data-field="F"/></div>`, "div")
	var p3 PageSpec
	if !dispatchTopLevelElement(actEl, &p3) {
		t.Errorf("dispatchTopLevelElement(action) = false")
	}
	// static-with-data path.
	staticData := firstElementNode(t, `<div><span data-bind="b"></span></div>`, "div")
	var p4 PageSpec
	if !dispatchTopLevelElement(staticData, &p4) {
		t.Errorf("dispatchTopLevelElement(static-data) = false")
	}
	// plain static path.
	staticEl := firstElementNode(t, `<h2>Title</h2>`, "h2")
	var p5 PageSpec
	if !dispatchTopLevelElement(staticEl, &p5) {
		t.Errorf("dispatchTopLevelElement(static) = false")
	}
}
