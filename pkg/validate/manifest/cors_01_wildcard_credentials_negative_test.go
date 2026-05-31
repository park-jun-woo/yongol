//ff:func feature=validate type=test control=sequence topic=manifest-cors
//ff:what CORS-01 테스트 — allow_origins=* + allow_credentials=true negative

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCors01WildcardCredentials_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				CORS: &pmanifest.CORSConfig{
					Enabled:          true,
					AllowOrigins:     []string{"*"},
					AllowCredentials: true,
				},
			},
		},
	}
	got := cors01WildcardCredentials(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[CORS-01]") {
		t.Fatalf("message missing [CORS-01] prefix: %q", got[0].Message)
	}
}
