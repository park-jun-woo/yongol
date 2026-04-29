//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what Run — manifest 검증 전체 실행 (C-*). parse 오류는 CLI 레벨에서 이미 게이팅됨.
package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all manifest validation rules. Parser-phase diagnostics are
// handled at the CLI boundary (cmd/yongol/validate_cmd.go); by the time this
// Run executes, fs.Manifest is guaranteed non-nil when KindConfig is present.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, c02APIVersion(fs)...)
	diags = append(diags, c03Kind(fs)...)
	diags = append(diags, c04MetadataName(fs)...)
	diags = append(diags, c05BackendModule(fs)...)
	diags = append(diags, c06BackendAuthRequired(fs)...)
	diags = append(diags, cors01WildcardCredentials(fs)...)
	diags = append(diags, obs01MetricsPath(fs)...)
	diags = append(diags, obs02MetricsPathNotOpenAPI(fs)...)
	diags = append(diags, obs03TracingExporter(fs)...)
	diags = append(diags, obs04TracingSampleRate(fs)...)
	diags = append(diags, sec201CookieWithoutCsrf(fs)...)
	diags = append(diags, sec301CspPermissive(fs)...)
	diags = append(diags, sec302HSTSShort(fs)...)
	diags = append(diags, sec401JWTSecretEnvRequired(fs)...)
	diags = append(diags, sec402AccessTTLUpperBound(fs)...)
	diags = append(diags, sec403AuthModeEnum(fs)...)
	// Phase004 (ssac/purify) — manifest-driven DB requirement checks.
	diags = append(diags, xnc90CacheBackendRequiresSQLC(fs)...)
	diags = append(diags, xns90SessionBackendRequiresSQLC(fs)...)
	diags = append(diags, xnq90QueueBackendRequiresSQLC(fs)...)
	diags = append(diags, xna90RefreshRequiresSQLC(fs)...)
	return diags
}
