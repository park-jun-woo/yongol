//ff:func feature=gen-gogin type=test control=sequence
//ff:what buildDomainCORSDispatch — 2도메인 prefix 분기/상속 + 1도메인 default 분기 + import 산출 검증

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildDomainCORSDispatch(t *testing.T) {
	backend := &manifest.CORSConfig{AllowOrigins: []string{"https://app.example.com"}}

	t.Run("multi domain orders specific-first and inherits backend", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{CORS: backend},
			Domains: map[string]manifest.DomainConfig{
				"public": {RoutePrefix: "/api"}, // inherits backend origins
				"admin":  {RoutePrefix: "/api/admin", CORS: &manifest.CORSConfig{AllowOrigins: []string{"https://admin.example.com"}}},
			},
		}}
		src, imports := buildDomainCORSDispatch(fs)
		// /api/admin (longer) must appear as an explicit case BEFORE the default.
		caseIdx := strings.Index(src, `case strings.HasPrefix(path, "/api/admin")`)
		defIdx := strings.Index(src, "default:")
		if caseIdx < 0 || defIdx < 0 || caseIdx > defIdx {
			t.Fatalf("admin case must precede default:\n%s", src)
		}
		if !strings.Contains(src, `slices.Contains([]string{"https://admin.example.com"}, origin)`) {
			t.Errorf("admin origins missing:\n%s", src)
		}
		if !strings.Contains(src, `slices.Contains([]string{"https://app.example.com"}, origin)`) {
			t.Errorf("public inherited origins missing:\n%s", src)
		}
		joined := strings.Join(imports, ",")
		if !strings.Contains(joined, `"slices"`) || !strings.Contains(joined, `"strings"`) {
			t.Errorf("multi-domain imports must include slices+strings, got %v", imports)
		}
	})

	t.Run("single domain has no switch and no strings import", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{CORS: backend},
			Domains: map[string]manifest.DomainConfig{
				"public": {RoutePrefix: "/api"},
			},
		}}
		src, imports := buildDomainCORSDispatch(fs)
		if strings.Contains(src, "switch") || strings.Contains(src, "strings.HasPrefix") {
			t.Errorf("single domain must not branch on path:\n%s", src)
		}
		if !strings.Contains(src, `return slices.Contains([]string{"https://app.example.com"}, origin)`) {
			t.Errorf("single domain origins missing:\n%s", src)
		}
		for _, imp := range imports {
			if imp == `"strings"` {
				t.Errorf("single domain must not import strings, got %v", imports)
			}
		}
	})
}
