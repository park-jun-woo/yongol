//ff:type feature=projectconfig type=model
//ff:what CSPConfig — backend.security_headers.csp 모델 (Content-Security-Policy 설정)

package manifest

// CSPConfig controls Content-Security-Policy. When Enabled is nil it is
// treated as true. ReportOnly swaps the emitted header name to
// Content-Security-Policy-Report-Only — directives still apply but the
// browser only logs violations instead of blocking resources.
//
// Directives maps CSP directive names (default-src, script-src, frame-ancestors,
// ...) to their source lists. Yongol concatenates the map into a single
// header string at codegen + runtime boot time; iteration order is made
// deterministic via sorted keys so generated output is reproducible.
type CSPConfig struct {
	Enabled    *bool               `yaml:"enabled,omitempty"`
	ReportOnly bool                `yaml:"report_only,omitempty"`
	Directives map[string][]string `yaml:"directives,omitempty"`
}
