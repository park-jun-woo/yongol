//ff:func feature=gen-gogin type=generator control=sequence
//ff:what flattenManifest — ssacmeta.EvaluateWhen 이 참조하는 map[string]any 조립

package infra

import "github.com/park-jun-woo/yongol/pkg/yongol"

// flattenManifest builds the map[string]any that ssacmeta.EvaluateWhen
// consults. Paths referenced by interface.yaml `when:` expressions must match
// the shape produced here.
//
// Supported paths today:
//
//	manifest.cache.backend
//	manifest.session.backend
//	manifest.queue.backend
//	manifest.backend.auth.refresh.enabled
//
// When fs.Manifest is nil we return an empty map — every `when:` on a
// non-"always" port resolves to false so nothing is emitted. This matches
// the expected behavior for memory-only dev runs.
func flattenManifest(fs *yongol.Fullstack) map[string]any {
	m := map[string]any{}
	if fs == nil || fs.Manifest == nil {
		return m
	}
	cfg := fs.Manifest
	// Infrastructure backends (cache/session/queue) live at the top level
	// of manifest.yaml — they are siblings of `backend:`, not children.
	if cfg.Cache != nil {
		m["cache"] = map[string]any{"backend": cfg.Cache.Backend}
	}
	if cfg.Session != nil {
		m["session"] = map[string]any{"backend": cfg.Session.Backend}
	}
	if cfg.Queue != nil {
		m["queue"] = map[string]any{"backend": cfg.Queue.Backend}
	}
	// backend.auth.refresh.enabled — auth package is the only DB-using
	// subsystem under `backend:`, and refresh_tokens support is implicit
	// whenever `backend.auth` is declared.
	backend := map[string]any{}
	if cfg.Backend.Auth != nil {
		backend["auth"] = map[string]any{
			"refresh": map[string]any{"enabled": true},
		}
	}
	m["backend"] = backend
	return m
}
