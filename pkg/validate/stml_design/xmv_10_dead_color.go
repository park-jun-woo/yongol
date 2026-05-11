//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XMV-10 — DESIGN.md colors 토큰이 STML에서 미참조 (WARNING)
package stml_design

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmv10DeadColor detects DESIGN.md color tokens not referenced by any STML page.
func xmv10DeadColor(fs *yongol.Fullstack, tokens pageTokenRefs) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Colors) == 0 {
		return nil
	}

	used := make(map[string]bool)
	for _, ref := range tokens.Colors {
		used[ref.Name] = true
	}

	var diags []diagnostic.Diagnostic
	names := sortedKeys(fs.DesignSpec.Colors)
	for _, name := range names {
		if !used[name] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.DesignSpec.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XMV-10] color token %q is defined in DESIGN.md but not referenced in any STML page", name),
				Advice:  fmt.Sprintf("Use %q in an STML class (e.g. bg-%s, text-%s) or remove it from DESIGN.md if unused", name, name, name),
			})
		}
	}
	return diags
}

// sortedKeys returns the keys of a map[string]string in sorted order.
func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
