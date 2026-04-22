//ff:type feature=report type=model topic=json
//ff:what yongol validate --format json 의 bespoke flat 스키마 (snake_case JSON 태그)
package json

// Document is the top-level yongol JSON report.
//
// Deliberately shallow and snake_case to differentiate from SARIF (which is
// camelCase + deeply nested). Intended for script / AI agent consumers that
// want `jq .summary.errors` to work without learning the SARIF data model.
type Document struct {
	YongolVersion string       `json:"yongol_version"`
	SpecsDir       string       `json:"specs_dir"`
	Summary        Summary      `json:"summary"`
	Diagnostics    []Diagnostic `json:"diagnostics"`
}

// Summary aggregates counts across the whole run.
//
// Checks is the number of validation rules executed in this run (i.e. the
// embedded catalog size, not the rules that fired). Consumers can use it as
// a sanity check — "0 checks" means the catalog never loaded.
type Summary struct {
	Errors   int `json:"errors"`
	Warnings int `json:"warnings"`
	Checks   int `json:"checks"`
}

// Diagnostic is a single finding (ERROR or WARNING). File / Line / Col may
// be empty / zero when the source position is unknown — consumers decide
// whether to surface that as "global" or suppress.
//
// Level is the internal Go enum value (`"ERROR"` / `"WARNING"`), NOT the
// SARIF lowercase variant. This is intentional.
type Diagnostic struct {
	RuleID  string `json:"rule_id"`
	Level   string `json:"level"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Col     int    `json:"col"`
	Message string `json:"message"`
}
