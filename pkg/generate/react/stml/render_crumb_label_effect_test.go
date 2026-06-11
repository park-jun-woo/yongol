//ff:func feature=stml-gen type=test control=sequence
//ff:what renderCrumbLabelEffect — null 가드 destructure·옵셔널 호출·title 갱신·suffix 유무 스냅샷 검증

package stml

import (
	"strings"
	"testing"
)

func TestRenderCrumbLabelEffect(t *testing.T) {
	t.Run("full snapshot with title suffix", func(t *testing.T) {
		want := "  const { setCrumbLabel } = useOutletContext<{ setCrumbLabel?: (label: string) => void }>() ?? {}\n" +
			"  useEffect(() => {\n" +
			"    const v = getBuildingData?.building_name\n" +
			"    if (v != null) {\n" +
			"      setCrumbLabel?.(String(v))\n" +
			"      document.title = String(v) + ' · zenflow'\n" +
			"    }\n" +
			"  }, [getBuildingData])\n\n"
		if got := renderCrumbLabelEffect("building_name", "getBuildingData", " · zenflow"); got != want {
			t.Errorf("got:\n%s\nwant:\n%s", got, want)
		}
	})

	t.Run("no suffix updates the bare label", func(t *testing.T) {
		got := renderCrumbLabelEffect("name", "getToolData", "")
		if want := "      document.title = String(v)\n"; !strings.Contains(got, want) {
			t.Errorf("got:\n%s\nmissing %q", got, want)
		}
	})
}
