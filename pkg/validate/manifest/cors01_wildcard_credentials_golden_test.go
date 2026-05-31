//ff:func feature=validate type=test control=sequence topic=manifest-cors
//ff:what cors01WildcardCredentials — allow_origins=* + credentials=true 금지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCors01WildcardCredentials_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				CORS: &pm.CORSConfig{
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
