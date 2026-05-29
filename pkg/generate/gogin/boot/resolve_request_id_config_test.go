//ff:func feature=gen-gogin type=test control=sequence topic=request-id
//ff:what resolveRequestIDConfig — trust_upstream + header 기본값 (true, "X-Request-Id")

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveRequestIDConfig(t *testing.T) {
	t.Run("defaults when absent", func(t *testing.T) {
		for _, fs := range []*yongol.Fullstack{
			nil,
			{},
			{Manifest: &pmanifest.ProjectConfig{}},
			{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Error: &pmanifest.ErrorConfig{}}}},
		} {
			trust, header := resolveRequestIDConfig(fs)
			if !trust || header != "X-Request-Id" {
				t.Errorf("expected (true, X-Request-Id), got (%v, %q)", trust, header)
			}
		}
	})

	t.Run("overrides applied", func(t *testing.T) {
		falseVal := false
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Error: &pmanifest.ErrorConfig{
				RequestID: &pmanifest.RequestIDConfig{TrustUpstream: &falseVal, Header: "X-Trace-Id"},
			}},
		}}
		trust, header := resolveRequestIDConfig(fs)
		if trust || header != "X-Trace-Id" {
			t.Errorf("expected (false, X-Trace-Id), got (%v, %q)", trust, header)
		}
	})
}
