//ff:type feature=rule type=model topic=catalog
//ff:what RuleMeta — rulebook.md 한 행(Rule ID / Level / Description / Source / Section) 메타
package catalog

// RuleMeta describes a single rule as declared in rulebook.md.
//
// SectionTitle is the H2 heading (e.g. "A. SSaC Internal") that the row belongs to.
// SectionAnchor is the GFM-style lowercase anchor of the section (e.g. "a-ssac-internal"),
// used by the SARIF emitter to build a stable helpUri.
type RuleMeta struct {
	ID            string
	Level         string // "ERROR" or "WARNING"
	Description   string
	Source        string // e.g. "pkg/validate/ssac/s_27_var_declared.go"
	SectionTitle  string
	SectionAnchor string
}
