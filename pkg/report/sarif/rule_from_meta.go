//ff:func feature=report type=util control=selection topic=sarif
//ff:what ruleFromMeta — catalog.RuleMeta → SARIF Rule 변환 (shortDescription/helpUri/defaultConfiguration)
package sarif

import (
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

// helpURIBase is the GitHub blob URL prefix used to link each rule back to
// its section in the canonical rulebook.md.
const helpURIBase = "https://github.com/park-jun-woo/yongol/blob/master/rulebook.md"

// ruleFromMeta renders a catalog RuleMeta as a SARIF Rule entry.
func ruleFromMeta(m rulecatalog.RuleMeta) Rule {
	r := Rule{
		ID:   m.ID,
		Name: m.ID, // canonical rulebook uses ID as the display name today
	}
	if m.Description != "" {
		r.ShortDescription = &Message{Text: m.Description}
	}
	if m.SectionAnchor != "" {
		r.HelpURI = helpURIBase + "#" + m.SectionAnchor
	}
	switch m.Level {
	case "ERROR":
		r.DefaultConfiguration = &DefaultConfiguration{Level: "error"}
	case "WARNING":
		r.DefaultConfiguration = &DefaultConfiguration{Level: "warning"}
	}
	if m.Source != "" {
		r.Properties = map[string]interface{}{"source": m.Source}
	}
	return r
}
