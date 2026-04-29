//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-6 — manifest.backend.auth 블록 부재 ERROR (yongol = auth 필수)

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c06BackendAuthRequired enforces yongol's project-level stance that every
// backend declared via manifest.yaml carries an auth block. yongol targets
// SaaS / business backends; auth-free dynamic backends are an anti-pattern
// (use a static site generator + CDN such as Hugo / Jekyll / Next.js SSG
// instead). When `backend.auth` is missing or nil the generator has no
// principal to inject into the request context, no claims to wire into
// `currentUser`, and no enforcement target for `@auth` annotations — every
// downstream stage degrades silently. Failing fast at validate time keeps
// SSOT authors from discovering the gap at generate / build time.
func c06BackendAuthRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Auth != nil {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:  "manifest.yaml",
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: "[C-6] manifest.backend.auth is required",
		Advice: "yongol does not support auth-free backends. Add a backend.auth " +
			"section (type: jwt + secret_env + claims) to manifest.yaml. " +
			"Public dynamic content should use a static site generator + CDN " +
			"(Hugo / Jekyll / Next.js SSG) instead of yongol.",
	}}
}
