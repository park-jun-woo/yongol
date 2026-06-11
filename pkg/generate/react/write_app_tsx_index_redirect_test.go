//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX "/" 인덱스 — 파일명 정렬 순 첫 공개 페이지로 redirect 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_IndexRedirect_FirstPublicPage(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "settings", FileName: "settings.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "about", FileName: "about.html"},
	}
	// about is protected → dashboard is the first public page by file name
	protected := map[string]bool{"about.html": true}
	if err := writeAppTSX(dir, pages, nil, "", protected, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, `<Route path="/" element={<Navigate to="/dashboard" replace />} />`)
	assertContains(t, content, `<Route path="*" element={<Navigate to="/" replace />} />`)
}
