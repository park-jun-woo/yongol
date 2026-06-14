//ff:func feature=stml-gen type=test control=sequence
//ff:what delete detail 페이지: 자기 GET removeQueries + 형제 목록 invalidate + navigate 결합 검증 (BUG-132 132-2)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_Delete_RemoveQueriesAndNavigate(t *testing.T) {
	page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetBuilding" data-param-buildingId="route.BuildingID">
    <h2 data-bind="building.name"></h2>
  </article>
  <section data-fetch="ListBuildingPhotos" data-param-buildingId="route.BuildingID">
    <ul data-each="photos"><li><span data-bind="caption"></span></li></ul>
  </section>
  <button data-action="DeleteBuilding" data-param-buildingId="route.BuildingID" data-redirect="/buildings">Delete</button>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		PathParamTypes: map[string]map[string]string{
			"DeleteBuilding": {"buildingId": "integer"},
			"GetBuilding":    {"buildingId": "integer"},
		},
	})

	// the deleted item's own GET is dropped from the cache (not refetched)
	assertContains(t, code, "queryClient.removeQueries({ queryKey: ['GetBuilding'] })")
	// the sibling list query is still invalidated
	assertContains(t, code, "queryClient.invalidateQueries({ queryKey: ['ListBuildingPhotos'] })")
	// the self GET is never invalidated (would refetch a 404)
	assertNotContains(t, code, "invalidateQueries({ queryKey: ['GetBuilding'] })")
	// and the screen navigates away
	assertContains(t, code, "navigate('/buildings')")
	assertContains(t, code, "const navigate = useNavigate()")
	assertContains(t, code, "useQueryClient")
}
