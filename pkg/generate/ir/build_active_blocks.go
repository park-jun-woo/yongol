//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildActiveBlocks -- 25개 부트 블록의 활성 상태를 정규화된 목록으로 조립

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildActiveBlocks evaluates all 25 blocks in the canonical order
// defined by baseCandidateBlocks and returns the full list with
// Active pre-evaluated.
func buildActiveBlocks(fs *yongol.Fullstack, ps *prepared.State) []BootBlock {
	hasAuthFlag := fs.Manifest != nil && fs.Manifest.Backend.Auth != nil && len(fs.Manifest.Backend.Auth.Claims) > 0
	hasAuthSeq := hasAuthSequence(fs)
	otelActive := otelEnabled(fs)
	corsActive := corsIsEnabled(fs)
	promActive := prometheusEnabled(fs)
	secHeadersActive := securityHeadersEnabled(fs)
	csrfActive := csrfIsActive(ps)
	rateLimitActive := rateLimitHasEntries(fs)

	return []BootBlock{
		{Name: "logger", Active: true},
		{Name: "env-helpers", Active: true},
		{Name: "db-init", Active: true},
		{Name: "jwt-secret", Active: hasAuthFlag},
		{Name: "authz-init", Active: hasAuthSeq},
		{Name: "session", Active: ps.ActiveBackends.Session != nil},
		{Name: "cache", Active: ps.ActiveBackends.Cache != nil},
		{Name: "file", Active: ps.ActiveBackends.File != nil},
		{Name: "server-struct", Active: true},
		{Name: "queue", Active: ps.ActiveBackends.Queue != nil},
		{Name: "otel-init", Active: otelActive},
		{Name: "router", Active: true},
		{Name: "request-id", Active: true},
		{Name: "error-envelope", Active: true},
		{Name: "cors", Active: corsActive},
		{Name: "prometheus", Active: promActive},
		{Name: "security-headers", Active: secHeadersActive},
		{Name: "csrf", Active: csrfActive},
		{Name: "body-limit", Active: true},
		{Name: "rate-limit", Active: rateLimitActive},
		{Name: "request-validator", Active: true},
		{Name: "health", Active: true},
		{Name: "register-handlers", Active: true},
		{Name: "auth-init", Active: ps.Auth.Present},
		{Name: "gin-run", Active: true},
	}
}
