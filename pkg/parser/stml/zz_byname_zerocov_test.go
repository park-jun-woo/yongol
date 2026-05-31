//ff:func feature=stml-parse type=test control=sequence
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package stml

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestByNameWalkTopLevel_ZeroCov exercises walkTopLevel / dispatchTopLevelElement
// and the static-with-data path by name.
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

// TestByNameParseBlocks_ZeroCov exercises parseFetchBlock / parseActionBlock /
// parseEachBlock / parseStateBind / parseStaticElement and their walkers.
func TestByNameParseBlocks_ZeroCov(t *testing.T) {
	fetchEl := firstElementNode(t,
		`<section data-fetch="List" class="c" data-param-id="x">
		   <span data-bind="title"></span>
		   <ul data-each="items"><li data-bind="name"></li></ul>
		   <p data-state="empty">none</p>
		   <span data-component="Card"></span>
		   <section data-fetch="Sub"></section>
		   <h3>Header</h3>
		 </section>`, "section")
	fb := parseFetchBlock(fetchEl, "List")
	if fb.OperationID != "List" {
		t.Errorf("parseFetchBlock OperationID = %q", fb.OperationID)
	}

	actEl := firstElementNode(t,
		`<div data-action="Create"><input data-field="Name"/><span data-component="Pick"></span><h4>x</h4><button type="submit">Go</button></div>`, "div")
	ab := parseActionBlock(actEl, "Create")
	if ab.OperationID != "Create" {
		t.Errorf("parseActionBlock OperationID = %q", ab.OperationID)
	}

	eachEl := firstElementNode(t,
		`<ul data-each="rows"><li data-bind="name"></li></ul>`, "ul")
	eb := parseEachBlock(eachEl, "rows")
	if eb.Field != "rows" {
		t.Errorf("parseEachBlock Field = %q", eb.Field)
	}

	stateEl := firstElementNode(t,
		`<p data-state="empty">없음<button data-action="Add">add</button><span>txt</span></p>`, "p")
	sb := parseStateBind(stateEl, "empty")
	if sb.Condition != "empty" {
		t.Errorf("parseStateBind Condition = %q", sb.Condition)
	}

	staticEl := firstElementNode(t, `<header class="h">Title<span>sub</span></header>`, "header")
	se := parseStaticElement(staticEl)
	if se.Tag != "header" {
		t.Errorf("parseStaticElement Tag = %q", se.Tag)
	}
}

// TestByNameWalkers_ZeroCov calls the walk* child-walkers by name.
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

// TestByNameDispatchChildren_ZeroCov calls dispatch*Child helpers by name.
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

// TestByNameDispatchStatic_ZeroCov exercises dispatchStatic* helpers by name.
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

// TestByNameHandlers_ZeroCov exercises handle* helpers by name.
func TestByNameHandlers_ZeroCov(t *testing.T) {
	compEl := firstElementNode(t, `<span data-component="Card" data-bind="b" data-field="f"></span>`, "span")
	var ab ActionBlock
	handleActionComponent(compEl, &ab)
	handleStaticActionComponent(compEl, &ab)

	fieldEl := firstElementNode(t, `<input data-field="Name" type="text"/>`, "input")
	handleActionField(fieldEl, &ab)
	handleStaticActionField(fieldEl, &ab)

	var fb FetchBlock
	handleFetchComponent(compEl, &fb)

	// walk-static action helpers
	se := StaticElement{Tag: "div"}
	var ab2 ActionBlock
	handleWalkStaticActionComponent(compEl, &ab2, &se)
	handleWalkStaticActionField(fieldEl, &ab2, &se)
}

// TestByNameStaticInContainers_ZeroCov exercises parseStaticIn* and
// walkStaticAction* helpers by name.
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

// TestByNameDiagCollectors_ZeroCov exercises the diag-collection helpers by name.
func TestByNameDiagCollectors_ZeroCov(t *testing.T) {
	eb := EachBlock{Diags: []diagnostic.Diagnostic{{Message: "x"}, {File: "f.html", Message: "y"}}}
	var out []diagnostic.Diagnostic
	appendEachDiags(&eb, "file.html", &out)
	if len(out) != 2 {
		t.Errorf("appendEachDiags out = %d, want 2", len(out))
	}

	fb := FetchBlock{
		Eaches:        []EachBlock{eb},
		NestedFetches: []FetchBlock{{Eaches: []EachBlock{eb}}},
	}
	var out2 []diagnostic.Diagnostic
	collectFetchDiags(&fb, "file.html", &out2)
	if len(out2) == 0 {
		t.Errorf("collectFetchDiags produced none")
	}

	cn := ChildNode{Fetch: &fb, Each: &eb}
	var out3 []diagnostic.Diagnostic
	collectChildDiags(&cn, "file.html", &out3)
	if len(out3) == 0 {
		t.Errorf("collectChildDiags produced none")
	}

	page := PageSpec{
		Fetches:  []FetchBlock{fb},
		Children: []ChildNode{cn},
	}
	d := collectEachDiags(&page, "file.html")
	if len(d) == 0 {
		t.Errorf("collectEachDiags produced none")
	}
}

// TestByNameLayout_ZeroCov exercises collectLayoutElement / walkLayoutNode by name.
func TestByNameLayout_ZeroCov(t *testing.T) {
	navEl := firstElementNode(t, `<a data-nav="/home">Home</a>`, "a")
	var layout LayoutSpec
	collectLayoutElement(navEl, &layout)
	if len(layout.NavItems) != 1 {
		t.Errorf("collectLayoutElement NavItems = %d, want 1", len(layout.NavItems))
	}

	slotEl := firstElementNode(t, `<slot data-outlet></slot>`, "slot")
	collectLayoutElement(slotEl, &layout)
	if !layout.HasOutlet {
		t.Errorf("collectLayoutElement HasOutlet = false")
	}

	nav := firstElementNode(t,
		`<nav><a data-nav="/a">A</a><a data-nav="/b">B</a><slot data-outlet></slot></nav>`, "nav")
	var layout2 LayoutSpec
	walkLayoutNode(nav, &layout2)
	if len(layout2.NavItems) != 2 {
		t.Errorf("walkLayoutNode NavItems = %d, want 2", len(layout2.NavItems))
	}
}

// firstTag returns the tag name of the first element in a simple fragment like
// "<tag ...>". Test-only helper.
func firstTag(frag string) string {
	i := 1 // skip '<'
	j := i
	for j < len(frag) {
		c := frag[j]
		if c == ' ' || c == '>' || c == '/' {
			break
		}
		j++
	}
	return frag[i:j]
}
