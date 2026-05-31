//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestModulePath(t *testing.T) {
	if got := modulePath(nil); got != "" {
		t.Errorf("nil = %q", got)
	}
	if got := modulePath(&yongol.Fullstack{}); got != "" {
		t.Errorf("nil manifest = %q", got)
	}
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Module: "github.com/x/y"}}}
	if got := modulePath(fs); got != "github.com/x/y" {
		t.Errorf("modulePath = %q", got)
	}
}
