//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestProjectID(t *testing.T) {
	if got := projectID(nil); got != "" {
		t.Errorf("nil = %q", got)
	}
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Metadata: manifest.Metadata{Name: "myapp"}}}
	if got := projectID(fs); got != "myapp" {
		t.Errorf("projectID = %q", got)
	}
}
