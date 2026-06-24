//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockCORS(domained) — AllowOrigins 제거 + AllowOriginWithContextFunc 설정 + gin/strings/slices import 검증

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockCORS_Domained(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Backend: manifest.Backend{CORS: &manifest.CORSConfig{
			Enabled:      true,
			AllowOrigins: []string{"https://app.example.com"},
			AllowMethods: []string{"GET", "POST"},
		}},
		Domains: map[string]manifest.DomainConfig{
			"public": {RoutePrefix: "/api"},
			"admin":  {RoutePrefix: "/api/admin"},
		},
	}}
	block := blockCORS(fs)
	helper := strings.Join(block.Funcs, "\n")
	if strings.Contains(helper, "cfg.AllowOrigins") || strings.Contains(helper, "CORS_ALLOW_ORIGINS") {
		t.Errorf("domain mode must NOT set AllowOrigins:\n%s", helper)
	}
	if !strings.Contains(helper, "AllowOriginWithContextFunc: func(c *gin.Context, origin string) bool {") {
		t.Errorf("domain mode must set AllowOriginWithContextFunc:\n%s", helper)
	}
	if !strings.Contains(helper, "isDomainOriginAllowed(origin, c.Request.URL.Path)") {
		t.Errorf("must dispatch to isDomainOriginAllowed:\n%s", helper)
	}
	imp := strings.Join(block.Imports, "\n")
	for _, must := range []string{`"github.com/gin-contrib/cors"`, `"github.com/gin-gonic/gin"`, `"slices"`, `"strings"`} {
		if !strings.Contains(imp, must) {
			t.Errorf("missing import %q in:\n%s", must, imp)
		}
	}
}
