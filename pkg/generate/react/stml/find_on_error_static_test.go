//ff:func feature=stml-gen type=test control=sequence
//ff:what findOnErrorStatic — data-on-error StaticElement 탐색(직접/재귀/부재) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFindOnErrorStatic(t *testing.T) {
	// direct hit: a static element carrying the marker, after a skipped
	// non-static node and a static node with nil Static.
	target := &stmlparser.StaticElement{Tag: "p", OnError: true}
	direct := []stmlparser.ChildNode{
		{Kind: "fetch", Fetch: &stmlparser.FetchBlock{}}, // not static -> skip
		{Kind: "static", Static: nil},                    // nil Static -> skip
		{Kind: "static", Static: target},
	}
	if got := findOnErrorStatic(direct); got != target {
		t.Errorf("direct = %+v, want %+v", got, target)
	}

	// recursive hit: marker nested inside a non-marker static element.
	nested := []stmlparser.ChildNode{
		{Kind: "static", Static: &stmlparser.StaticElement{
			Tag: "div",
			Children: []stmlparser.ChildNode{
				{Kind: "static", Static: target},
			},
		}},
	}
	if got := findOnErrorStatic(nested); got != target {
		t.Errorf("recursive = %+v, want %+v", got, target)
	}

	// absent: no marker anywhere -> nil
	none := []stmlparser.ChildNode{
		{Kind: "static", Static: &stmlparser.StaticElement{Tag: "span"}},
	}
	if got := findOnErrorStatic(none); got != nil {
		t.Errorf("absent = %+v, want nil", got)
	}
}
