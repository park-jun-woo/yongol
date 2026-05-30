//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func boolPtr(b bool) *bool { return &b }

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

func TestProjectID(t *testing.T) {
	if got := projectID(nil); got != "" {
		t.Errorf("nil = %q", got)
	}
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Metadata: manifest.Metadata{Name: "myapp"}}}
	if got := projectID(fs); got != "myapp" {
		t.Errorf("projectID = %q", got)
	}
}

func TestRateLimitHasEntries(t *testing.T) {
	if rateLimitHasEntries(nil) {
		t.Errorf("nil should be false")
	}
	empty := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if rateLimitHasEntries(empty) {
		t.Errorf("empty should be false")
	}
	withRL := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		RateLimit: manifest.RateLimitConfig{"op": {Rate: 10, Period: "1m"}},
	}}}
	if !rateLimitHasEntries(withRL) {
		t.Errorf("with entry should be true")
	}
}

func TestCorsIsEnabled(t *testing.T) {
	if corsIsEnabled(nil) {
		t.Errorf("nil should be false")
	}
	noCORS := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if corsIsEnabled(noCORS) {
		t.Errorf("nil CORS should be false")
	}
	on := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		CORS: &manifest.CORSConfig{Enabled: true},
	}}}
	if !corsIsEnabled(on) {
		t.Errorf("enabled CORS should be true")
	}
}

func TestOtelEnabled(t *testing.T) {
	if otelEnabled(nil) {
		t.Errorf("nil should be false")
	}
	off := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if otelEnabled(off) {
		t.Errorf("no observability should be false")
	}
	on := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		Observability: &manifest.Observability{Tracing: &manifest.ObservabilityTracing{Enabled: true}},
	}}}
	if !otelEnabled(on) {
		t.Errorf("tracing enabled should be true")
	}
}

func TestPrometheusEnabled(t *testing.T) {
	// defaults to true (opt-out)
	if !prometheusEnabled(nil) {
		t.Errorf("nil should default true")
	}
	if !prometheusEnabled(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}) {
		t.Errorf("no observability should default true")
	}
	disabled := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		Observability: &manifest.Observability{Metrics: &manifest.ObservabilityMetrics{Enabled: boolPtr(false)}},
	}}}
	if prometheusEnabled(disabled) {
		t.Errorf("explicit false should be false")
	}
}

func TestSecurityHeadersEnabled(t *testing.T) {
	if !securityHeadersEnabled(nil) {
		t.Errorf("nil should default true")
	}
	if !securityHeadersEnabled(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}) {
		t.Errorf("no config should default true")
	}
	disabled := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		SecurityHeaders: &manifest.SecurityHeadersConfig{Enabled: boolPtr(false)},
	}}}
	if securityHeadersEnabled(disabled) {
		t.Errorf("explicit false should be false")
	}
}

func TestHasAuthSequence(t *testing.T) {
	withAuth := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
		{Sequences: []ssac.Sequence{{Type: "get"}, {Type: "auth"}}},
	}}
	if !hasAuthSequence(withAuth) {
		t.Errorf("expected auth sequence found")
	}
	without := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{
		{Sequences: []ssac.Sequence{{Type: "get"}}},
	}}
	if hasAuthSequence(without) {
		t.Errorf("expected no auth sequence")
	}
}
