//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestParseLinkInFetch — fetch 블록 자식 data-link 파싱(bind 등록 + ChildNode link) 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseLinkInFetch(t *testing.T) {
	input := `<main>
  <section data-fetch="GetBuilding" data-param-building-id="route.BuildingID">
    <a data-link="unit-list" data-link-params="route.BuildingID -> BuildingID">
      <span data-bind="name"></span>
    </a>
  </section>
</main>`

	page, diags := ParseReader("building-detail.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if len(page.Fetches) != 1 {
		t.Fatalf("expected 1 fetch, got %d", len(page.Fetches))
	}
	fb := page.Fetches[0]

	var link *LinkRef
	for _, ch := range fb.Children {
		if ch.Kind == "link" {
			link = ch.Link
		}
	}
	if link == nil {
		t.Fatalf("no link ChildNode in fetch children: %+v", fb.Children)
	}
	if link.TargetPage != "unit-list" {
		t.Errorf("TargetPage = %q", link.TargetPage)
	}
	if len(link.Params) != 1 || link.Params[0].Source != "route.BuildingID" || link.Params[0].Segment != "BuildingID" {
		t.Errorf("Params = %+v", link.Params)
	}
	// The bind child stays under the link and registers with the fetch.
	if len(link.Children) != 1 || link.Children[0].Kind != "bind" {
		t.Errorf("link children = %+v", link.Children)
	}
	if len(fb.Binds) != 1 || fb.Binds[0].Name != "name" {
		t.Errorf("fb.Binds = %+v", fb.Binds)
	}
}
