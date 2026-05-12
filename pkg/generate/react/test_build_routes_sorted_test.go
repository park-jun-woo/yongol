//ff:func feature=gen-react type=test control=sequence
//ff:what buildRoutes 라우트 정렬 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildRoutes_Sorted(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "zebra", FileName: "zebra.html"},
		{Name: "alpha", FileName: "alpha.html"},
		{Name: "mid", FileName: "mid.html"},
	}
	routes := buildRoutes(pages, "")
	if len(routes) != 3 {
		t.Fatalf("got %d routes, want 3", len(routes))
	}
	if routes[0].Path != "/alpha" {
		t.Errorf("routes[0].Path = %q, want /alpha", routes[0].Path)
	}
	if routes[1].Path != "/mid" {
		t.Errorf("routes[1].Path = %q, want /mid", routes[1].Path)
	}
	if routes[2].Path != "/zebra" {
		t.Errorf("routes[2].Path = %q, want /zebra", routes[2].Path)
	}
}
