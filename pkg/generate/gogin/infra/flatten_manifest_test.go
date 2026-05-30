//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestFlattenManifest — manifest → when-eval map 조립 (nil/cache/session/queue/auth) 검증

package infra

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFlattenManifest(t *testing.T) {
	t.Run("NilManifestEmptyMap", func(t *testing.T) {
		m := flattenManifest(&yongol.Fullstack{})
		// backend key is always present but empty; no cache/session/queue.
		if _, ok := m["cache"]; ok {
			t.Errorf("expected no cache key, got: %v", m)
		}
	})

	t.Run("AllBackendsSet", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Cache:   &pmanifest.BuiltinBackend{Backend: "redis"},
				Session: &pmanifest.BuiltinBackend{Backend: "redis"},
				Queue:   &pmanifest.QueueBackend{Backend: "memory"},
				Backend: pmanifest.Backend{Auth: &pmanifest.Auth{Type: "jwt"}},
			},
		}
		m := flattenManifest(fs)

		if got := m["cache"].(map[string]any)["backend"]; got != "redis" {
			t.Errorf("cache.backend = %v, want redis", got)
		}
		if got := m["session"].(map[string]any)["backend"]; got != "redis" {
			t.Errorf("session.backend = %v, want redis", got)
		}
		if got := m["queue"].(map[string]any)["backend"]; got != "memory" {
			t.Errorf("queue.backend = %v, want memory", got)
		}
		backend := m["backend"].(map[string]any)
		auth := backend["auth"].(map[string]any)
		refresh := auth["refresh"].(map[string]any)
		if refresh["enabled"] != true {
			t.Errorf("backend.auth.refresh.enabled = %v, want true", refresh["enabled"])
		}
	})
}
