//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockCORS — gin-contrib/cors 미들웨어 등록 (manifest + env 기반)

package boot

import (
	"strings"
	"testing"
	"time"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockCORS_Disabled(t *testing.T) {
	block := blockCORS(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}})
	if block.Name != "cors" {
		t.Errorf("name = %q, want cors", block.Name)
	}
	if len(block.Lines) != 0 || len(block.Imports) != 0 || len(block.Funcs) != 0 {
		t.Fatalf("disabled CORS must be inert, got %+v", block)
	}
}

func TestBlockCORS_Enabled(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{CORS: &pmanifest.CORSConfig{
			Enabled:          true,
			AllowOrigins:     []string{"https://app.example.com"},
			AllowMethods:     []string{"GET", "POST"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}},
	}}
	block := blockCORS(fs)
	if !strings.Contains(strings.Join(block.Lines, "\n"), "r.Use(cors.New(buildCORSConfig()))") {
		t.Errorf("must register cors middleware, got:\n%v", block.Lines)
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `"github.com/gin-contrib/cors"`) {
		t.Errorf("must import gin-contrib/cors, got:\n%v", block.Imports)
	}
	helper := strings.Join(block.Funcs, "\n")
	for _, must := range []string{
		`envStringList("CORS_ALLOW_ORIGINS", []string{"https://app.example.com"})`,
		`envStringList("CORS_ALLOW_METHODS", []string{"GET", "POST"})`,
		`envBool("CORS_ALLOW_CREDENTIALS", true)`,
	} {
		if !strings.Contains(helper, must) {
			t.Errorf("helper missing %q, got:\n%s", must, helper)
		}
	}
}
