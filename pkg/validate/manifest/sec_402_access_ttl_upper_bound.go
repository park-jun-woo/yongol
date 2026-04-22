//ff:func feature=validate type=rule control=sequence topic=manifest-auth
//ff:what SEC-402 — backend.auth.access_token_ttl 은 30분 이하여야 함 (OWASP 권고)

package manifest

import (
	"time"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// accessTTLUpperBound is the recommended ceiling for access-token lifetime.
// 30 minutes matches OWASP Cheat Sheet guidance: a long-lived access token
// enlarges the blast radius of a leaked JWT because rotation via the
// refresh flow is the only revocation path. TTLs above this threshold are
// flagged as WARNING rather than ERROR so projects with an explicit risk
// acceptance can still ship, but the diagnostic is visible in CI.
var accessTTLUpperBound = 30 * time.Minute

// sec402AccessTTLUpperBound warns when manifest.backend.auth.access_token_ttl
// is parseable and exceeds 30m. Missing field or unparseable duration
// produce no diagnostic — other rules (or the default 15m) handle those.
func sec402AccessTTLUpperBound(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	raw := fs.Manifest.Backend.Auth.AccessTokenTTL
	if raw == "" {
		return nil
	}
	d, err := time.ParseDuration(raw)
	if err != nil {
		return nil
	}
	if d <= accessTTLUpperBound {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[SEC-402] backend.auth.access_token_ttl=" + raw + " 은 권장 상한(30m)을 초과합니다",
		Advice:  "access_token_ttl 을 30m 이하로 낮추고 refresh-token rotation 주기를 짧게 유지하세요",
	}}
}
