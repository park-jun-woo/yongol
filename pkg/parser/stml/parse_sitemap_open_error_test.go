//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemap_OpenErrorIsDiag — 존재하지 않는 sitemap 파일 경로는 파싱 에러 진단 검증

package stml

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseSitemap_OpenErrorIsDiag(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sitemap.html")
	spec, diags := ParseSitemap(path)
	if len(diags) != 1 {
		t.Fatalf("expected 1 open-error diag, got %+v", diags)
	}
	d := diags[0]
	if d.File != path || d.Phase != diagnostic.PhaseParse || d.Level != diagnostic.LevelError {
		t.Errorf("diag meta = %+v", d)
	}
	if d.Message == "" {
		t.Errorf("Message should carry the os.Open error")
	}
	if len(spec.Navs) != 0 || spec.FileName != "" {
		t.Errorf("spec = %+v, want zero value on error", spec)
	}
}
