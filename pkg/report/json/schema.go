//ff:type feature=report type=model topic=json
//ff:what Document — yongol validate --format json 의 bespoke flat 스키마 (top-level)
package json

// Document is the top-level yongol JSON report.
//
// Deliberately shallow and snake_case to differentiate from SARIF (which is
// camelCase + deeply nested). Intended for script / AI agent consumers that
// want `jq .summary.errors` to work without learning the SARIF data model.
type Document struct {
	YongolVersion string       `json:"yongol_version"`
	SpecsDir      string       `json:"specs_dir"`
	Summary       Summary      `json:"summary"`
	Diagnostics   []Diagnostic `json:"diagnostics"`
}
