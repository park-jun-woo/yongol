//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestParseEachRowAction — data-each 내 data-action 파싱 (직접 자식·static 래퍼, item.* 파라미터, 거부 진단 없음)

package stml

import (
	"strings"
	"testing"
)

func TestParseEachRowAction(t *testing.T) {
	input := `<main>
  <section data-fetch="GetUnit" data-param-building-id="route.BuildingID">
    <ul data-each="photos">
      <li>
        <span data-bind="caption"></span>
        <button data-action="DeletePhoto"
                data-param-building-id="route.BuildingID"
                data-param-photo-id="item.id">삭제</button>
        <div>
          <button data-action="StarPhoto" data-param-photo-id="item.id">대표</button>
        </div>
      </li>
    </ul>
  </section>
</main>`

	page, diags := ParseReader("unit-info.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatalf("row actions must parse without diagnostics, got %v", diags)
	}

	if len(page.Fetches) != 1 || len(page.Fetches[0].Eaches) != 1 {
		t.Fatalf("expected 1 fetch with 1 each, got %+v", page.Fetches)
	}
	each := page.Fetches[0].Eaches[0]

	// Both the direct child action and the static-wrapped action are
	// collected on the EachBlock.
	if len(each.Actions) != 2 {
		t.Fatalf("each.Actions = %d, want 2: %+v", len(each.Actions), each.Actions)
	}
	if each.Actions[0].OperationID != "DeletePhoto" || each.Actions[1].OperationID != "StarPhoto" {
		t.Errorf("unexpected action order: %+v", each.Actions)
	}

	// item.* and route.* param sources survive on the action.
	del := each.Actions[0]
	if len(del.Params) != 2 {
		t.Fatalf("DeletePhoto params = %d, want 2: %+v", len(del.Params), del.Params)
	}
	bySource := map[string]string{}
	for _, p := range del.Params {
		bySource[p.Name] = p.Source
	}
	if bySource["buildingId"] != "route.BuildingID" || bySource["photoId"] != "item.id" {
		t.Errorf("unexpected param sources: %v", bySource)
	}

	// The direct child appears as a ChildNode{Kind: "action"} in DOM order.
	var kinds []string
	for _, ch := range each.Children {
		kinds = append(kinds, ch.Kind)
	}
	want := []string{"bind", "action", "static"}
	if strings.Join(kinds, ",") != strings.Join(want, ",") {
		t.Errorf("each.Children kinds = %v, want %v", kinds, want)
	}
	if each.Children[1].Action == nil || each.Children[1].Action.OperationID != "DeletePhoto" {
		t.Errorf("Children[1] is not the DeletePhoto action: %+v", each.Children[1])
	}
	// The static wrapper carries the nested action node.
	st := each.Children[2].Static
	if st == nil || len(st.Children) != 1 || st.Children[0].Kind != "action" {
		t.Fatalf("static wrapper does not carry the nested action: %+v", st)
	}
}
