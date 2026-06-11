//ff:func feature=gen-react type=test control=sequence
//ff:what fsSitemap — nil Fullstack/nil Sitemap/설정된 Sitemap 반환 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFsSitemap(t *testing.T) {
	if got := fsSitemap(nil); got != nil {
		t.Errorf("nil Fullstack must yield nil, got %v", got)
	}
	if got := fsSitemap(&yongol.Fullstack{}); got != nil {
		t.Errorf("absent sitemap must yield nil, got %v", got)
	}
	sm := &stml.SitemapSpec{FileName: "sitemap.html"}
	if got := fsSitemap(&yongol.Fullstack{Sitemap: sm}); got != sm {
		t.Errorf("expected the Fullstack's sitemap back, got %v", got)
	}
}
