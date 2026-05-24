//ff:func feature=validate type=test control=sequence topic=rego-manifest
//ff:what Run — Rego+Manifest 전체 검증 (빈 fs) 검증

package rego_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_RegoManifest(t *testing.T) {
	t.Run("empty fullstack returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
