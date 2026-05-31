//ff:func feature=stml-parse type=test control=sequence
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

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
