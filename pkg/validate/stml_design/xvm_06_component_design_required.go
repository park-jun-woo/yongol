//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XVM-06 — STML data-component가 DESIGN.md components에 정의되지 않으면 ERROR
package stml_design

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xvm06ComponentDesignRequired detects STML data-component references that
// have no corresponding entry in DESIGN.md components.
func xvm06ComponentDesignRequired(fs *yongol.Fullstack, tokens pageTokenRefs) []diagnostic.Diagnostic {
	if len(tokens.Components) == 0 {
		return nil
	}

	defined := make(map[string]bool)
	if fs.DesignSpec != nil {
		for name := range fs.DesignSpec.Components {
			defined[name] = true
		}
	}

	// Collect unique missing names with their first occurrence file.
	type miss struct {
		name string
		file string
	}
	seen := make(map[string]bool)
	var missing []miss
	for _, ref := range tokens.Components {
		if defined[ref.Name] {
			continue
		}
		if seen[ref.Name] {
			continue
		}
		seen[ref.Name] = true
		missing = append(missing, miss{name: ref.Name, file: ref.File})
	}

	// Sort by name for deterministic output.
	sort.Slice(missing, func(i, j int) bool {
		return missing[i].name < missing[j].name
	})

	var diags []diagnostic.Diagnostic
	for _, m := range missing {
		diags = append(diags, diagnostic.Diagnostic{
			File:    m.file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XVM-06] data-component %q is used in STML but not defined in DESIGN.md components", m.name),
			Advice:  fmt.Sprintf("Add a %q entry under components in DESIGN.md or remove the data-component attribute from STML", m.name),
		})
	}
	return diags
}
