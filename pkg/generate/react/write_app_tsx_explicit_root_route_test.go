//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 기존 "/" 라우트 보유 시 인덱스 redirect 미방출(충돌 없음) 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_ExplicitRootRoute_NoIndexRedirect(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "home", FileName: "home.html", Route: "/"},
		{Name: "about", FileName: "about.html"},
	}
	if err := writeAppTSX(dir, pages, nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// the page that claims "/" is the index — no redirect route on top of it
	assertContains(t, content, `<Route path="/" element={<Home />} />`)
	assertNotContains(t, content, `<Route path="/" element={<Navigate`)
	if strings.Count(content, `<Route path="/" `) != 1 {
		t.Errorf("expected exactly one \"/\" route, got %d", strings.Count(content, `<Route path="/" `))
	}
	// catch-all is still emitted
	assertContains(t, content, `<Route path="*" element={<Navigate to="/" replace />} />`)
}
