//ff:func feature=gen-react type=test control=sequence
//ff:what TestRenderSitemapGroupQueries — 복수 op 의 useQuery const 순서 방출·빈 목록 무방출 검증

package react

import (
	"strings"
	"testing"
)

func TestRenderSitemapGroupQueries(t *testing.T) {
	var sb strings.Builder
	renderSitemapGroupQueries(&sb, []string{"ListA", "ListB"}, false)
	got := sb.String()
	a, b := strings.Index(got, "listAData"), strings.Index(got, "listBData")
	if a == -1 || b == -1 || a > b {
		t.Errorf("queries must emit in document order:\n%s", got)
	}

	sb.Reset()
	renderSitemapGroupQueries(&sb, nil, true)
	if sb.Len() != 0 {
		t.Errorf("no ops must emit nothing, got:\n%s", sb.String())
	}
}
