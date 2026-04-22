//ff:func feature=validate type=test control=sequence topic=manifest-cors
//ff:what CORS-01 테스트 — allow_origins=* + allow_credentials=true golden

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestCors01WildcardCredentials_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				CORS: &pmanifest.CORSConfig{
					Enabled:          true,
					AllowOrigins:     []string{"https://a.com"},
					AllowCredentials: true,
				},
			},
		},
	}
	if got := cors01WildcardCredentials(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
