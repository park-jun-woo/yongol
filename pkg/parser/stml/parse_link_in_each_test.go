//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestParseLinkInEach — each 행 자식 data-link 파싱(bind 등록 + ChildNode link) 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseLinkInEach(t *testing.T) {
	input := `<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li>
        <a data-link="building-detail" data-link-params="item.id -> BuildingID">
          <span data-bind="name"></span>
        </a>
      </li>
    </ul>
  </section>
</main>`

	page, diags := ParseReader("building-list.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if len(page.Fetches) != 1 || len(page.Fetches[0].Eaches) != 1 {
		t.Fatalf("expected 1 fetch with 1 each, got %+v", page.Fetches)
	}
	eb := page.Fetches[0].Eaches[0]

	var link *LinkRef
	for _, ch := range eb.Children {
		if ch.Kind == "link" {
			link = ch.Link
		}
	}
	if link == nil {
		t.Fatalf("no link ChildNode in each children: %+v", eb.Children)
	}
	if link.TargetPage != "building-detail" {
		t.Errorf("TargetPage = %q", link.TargetPage)
	}
	if len(link.Params) != 1 || link.Params[0].Source != "item.id" {
		t.Errorf("Params = %+v", link.Params)
	}
	// The bind child stays under the link for codegen…
	if len(link.Children) != 1 || link.Children[0].Kind != "bind" || link.Children[0].Bind.Name != "name" {
		t.Errorf("link children = %+v", link.Children)
	}
	// …and registers with the enclosing each for validation (TM-06).
	if len(eb.Binds) != 1 || eb.Binds[0].Name != "name" {
		t.Errorf("eb.Binds = %+v", eb.Binds)
	}
}
