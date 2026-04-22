//ff:type feature=rule type=model topic=catalog
//ff:what Catalog — stores the parsed rulebook.md result as a Rule ID keyed lookup table
package catalog

import "strings"

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

// Catalog is an ordered slice of rules plus an id→index map for O(1) lookup.
// The slice order follows the rulebook.md order, which the SARIF emitter
// relies on for stable `ruleIndex` values within a run.
type Catalog struct {
	Rules   []RuleMeta
	byID    map[string]int
}

// NewCatalog builds a Catalog from a pre-parsed rule slice.
func NewCatalog(rules []RuleMeta) *Catalog {
	c := &Catalog{
		Rules: rules,
		byID:  make(map[string]int, len(rules)),
	}
	for i, r := range rules {
		c.byID[r.ID] = i
	}
	return c
}

// Lookup returns the RuleMeta for the given rule ID and whether it was found.
func (c *Catalog) Lookup(ruleID string) (RuleMeta, bool) {
	if c == nil {
		return RuleMeta{}, false
	}
	i, ok := c.byID[ruleID]
	if !ok {
		return RuleMeta{}, false
	}
	return c.Rules[i], true
}

// Index returns the 0-based index of the given rule ID in the Rules slice.
// Returns -1 when the rule is not present.
func (c *Catalog) Index(ruleID string) int {
	if c == nil {
		return -1
	}
	if i, ok := c.byID[ruleID]; ok {
		return i
	}
	return -1
}

// Len returns the number of rules in the catalog.
func (c *Catalog) Len() int {
	if c == nil {
		return 0
	}
	return len(c.Rules)
}

// sectionAnchor converts an H2 heading into a GFM-style slug:
// lowercase, dots/spaces → hyphen, punctuation stripped.
// Used for helpUri construction.
func sectionAnchor(title string) string {
	var b strings.Builder
	b.Grow(len(title))
	prevDash := false
	for _, r := range strings.ToLower(title) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			prevDash = false
		case r == ' ', r == '-', r == '.', r == '/':
			if !prevDash && b.Len() > 0 {
				b.WriteByte('-')
				prevDash = true
			}
		}
	}
	out := b.String()
	return strings.TrimRight(out, "-")
}
