//ff:func feature=stml-parse type=test control=sequence
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"
)

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
