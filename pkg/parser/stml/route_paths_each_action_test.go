//ff:func feature=stml-parse type=test control=sequence
//ff:what TestRoutePathsEachAction — each 내 행 액션의 route.* 는 선택 세그먼트로 참여, item.* 는 제외

package stml

import (
	"reflect"
	"strings"
	"testing"
)

func TestRoutePathsEachAction(t *testing.T) {
	input := `<main>
  <section data-fetch="GetUnit" data-param-building-id="route.BuildingID">
    <ul data-each="photos">
      <li>
        <span data-bind="caption"></span>
        <button data-action="DeletePhoto"
                data-param-building-id="route.BuildingID"
                data-param-photo-id="item.id">삭제</button>
      </li>
    </ul>
  </section>
</main>`
	page, diags := ParseReader("unit-info.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	// item.id contributes no segment; route.BuildingID is already required
	// via the fetch — the derived route has exactly one segment.
	if got := RoutePaths(page); !reflect.DeepEqual(got, []string{"/unit-info/:BuildingID"}) {
		t.Errorf("RoutePaths = %v, want [/unit-info/:BuildingID]", got)
	}

	// A route.* consumed only by the row action joins as an optional
	// trailing segment (same rule as any child action).
	input2 := `<main>
  <section data-fetch="ListPhotos">
    <ul data-each="photos">
      <li>
        <button data-action="MovePhoto"
                data-param-album-id="route.AlbumID"
                data-param-photo-id="item.id">이동</button>
      </li>
    </ul>
  </section>
</main>`
	page2, diags2 := ParseReader("photos.html", strings.NewReader(input2))
	if len(diags2) > 0 {
		t.Fatal(diags2)
	}
	if got := RoutePaths(page2); !reflect.DeepEqual(got, []string{"/photos/:AlbumID?"}) {
		t.Errorf("RoutePaths = %v, want [/photos/:AlbumID?]", got)
	}
}
