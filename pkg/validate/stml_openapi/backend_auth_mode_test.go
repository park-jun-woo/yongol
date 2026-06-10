//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what backendAuthMode — nil/manifest 없음/auth 없음은 "", 선언 시 ResolvedMode 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBackendAuthMode(t *testing.T) {
	t.Run("nil fullstack returns empty", func(t *testing.T) {
		if got := backendAuthMode(nil); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("nil manifest returns empty", func(t *testing.T) {
		if got := backendAuthMode(&yongol.Fullstack{}); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("no auth block returns empty", func(t *testing.T) {
		fs := makeFS(nil, nil)
		if got := backendAuthMode(fs); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("declared auth returns resolved mode", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "bearer")
		if got := backendAuthMode(fs); got != "bearer" {
			t.Errorf("expected bearer, got %q", got)
		}
	})
}
