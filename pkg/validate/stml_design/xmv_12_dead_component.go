//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XMV-12 — DESIGN.md components 토큰이 STML에서 미참조 (WARNING)
package stml_design

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmv12DeadComponent detects DESIGN.md component tokens not referenced by any
// STML data-component attribute.
func xmv12DeadComponent(fs *yongol.Fullstack, tokens pageTokenRefs) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Components) == 0 {
		return nil
	}

	used := make(map[string]bool)
	for _, ref := range tokens.Components {
		used[ref.Name] = true
	}

	var diags []diagnostic.Diagnostic
	names := sortedCompKeys(fs.DesignSpec.Components)
	for _, name := range names {
		if !used[name] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XMV-12] component token %q is defined in DESIGN.md but not referenced in any STML page", name),
				Advice:  fmt.Sprintf("Use data-component=%q in an STML page or remove it from DESIGN.md if unused", name),
			})
		}
	}
	return diags
}

// sortedCompKeys returns the keys of a ComponentToken map in sorted order.
func sortedCompKeys(m map[string]design.ComponentToken) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
