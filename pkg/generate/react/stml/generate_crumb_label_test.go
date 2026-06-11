//ff:func feature=stml-gen type=test control=sequence
//ff:what GeneratePage — CrumbFields 등재 페이지만 useOutletContext+라벨 effect 방출 / 미등재·fetch 부재 미배선 검증

package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_CrumbLabel(t *testing.T) {
	const src = `<main>
  <article data-fetch="GetBuilding" data-param-building-id="route.BuildingID">
    <span data-bind="building_name"></span>
  </article>
</main>`

	t.Run("declared page wires useOutletContext and the label effect", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(src))
		code := GeneratePage(page, "", GenerateOptions{
			DocumentTitles:   map[string]string{"building-detail": "건물 상세 · zenflow"},
			CrumbFields:      map[string]string{"building-detail": "building_name"},
			CrumbTitleSuffix: " · zenflow",
		})
		assertContains(t, code, "useOutletContext } from 'react-router-dom'")
		// the null guard: outlet context defaults to null on bare pages
		assertContains(t, code, "const { setCrumbLabel } = useOutletContext<{ setCrumbLabel?: (label: string) => void }>() ?? {}")
		assertContains(t, code, "    const v = getBuildingData?.building_name\n")
		assertContains(t, code, "      setCrumbLabel?.(String(v))\n")
		assertContains(t, code, "      document.title = String(v) + ' · zenflow'\n")
		assertContains(t, code, "  }, [getBuildingData])\n")
		// the static mount title stays — the pre-arrival fallback
		assertContains(t, code, "document.title = '건물 상세 · zenflow'")
	})

	t.Run("undeclared page wires no outlet context — Phase005 output stays byte-identical", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(src))
		code := GeneratePage(page, "", GenerateOptions{
			DocumentTitles: map[string]string{"building-detail": "건물 상세 · zenflow"},
		})
		assertNotContains(t, code, "useOutletContext")
		assertNotContains(t, code, "setCrumbLabel")
	})

	t.Run("declared page without a fetch wires nothing (TM-50 territory)", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("about.html", strings.NewReader(`<main><p>소개</p></main>`))
		code := GeneratePage(page, "", GenerateOptions{
			CrumbFields: map[string]string{"about": "name"},
		})
		assertNotContains(t, code, "useOutletContext")
		assertNotContains(t, code, "setCrumbLabel")
	})
}
