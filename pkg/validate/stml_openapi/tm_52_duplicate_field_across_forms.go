//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-52 — 한 페이지의 둘 이상 data-action 폼이 같은 data-field 이름을 선언 (WARNING, DOM id 폼 스코프 안내)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm52DuplicateFieldAcrossForms warns when two or more data-action forms on
// one page declare a same-named data-field. Before BUG-127's codegen fix the
// generated id/htmlFor were the bare field name, so such a page emitted a
// duplicate DOM id (label-for breakage). The id is now form-scoped to
// "{operationId}-{field}", so this is no longer a defect — the rule fills the
// detection gap and tells the user the bare field name is not a stable DOM id
// to target from external selectors (E2E/CSS). It collects every form actually
// rendered on the page via CollectChildActions(page.Children) — page.Actions
// holds only top-level forms and misses the nested update/create forms of the
// BUG-127 repro. data-component fields emit no id and are excluded.
func tm52DuplicateFieldAcrossForms(page stml.PageSpec) []diagnostic.Diagnostic {
	actions := stml.CollectChildActions(page.Children)
	if len(actions) < 2 {
		return nil
	}

	// field name -> operationIds of the forms declaring it (one entry per
	// form, so two same-operationId forms still register as a collision).
	declarers := map[string][]string{}
	order := []string{}
	for _, a := range actions {
		seen := map[string]bool{}
		for _, f := range a.Fields {
			if strings.HasPrefix(f.Tag, "data-component:") {
				continue
			}
			if seen[f.Name] {
				continue
			}
			seen[f.Name] = true
			if _, ok := declarers[f.Name]; !ok {
				order = append(order, f.Name)
			}
			declarers[f.Name] = append(declarers[f.Name], a.OperationID)
		}
	}

	var diags []diagnostic.Diagnostic
	for _, name := range order {
		ops := declarers[name]
		if len(ops) < 2 {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    page.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[TM-52] data-field %q is declared by %d forms on this page (operationIds: %s); their generated DOM id/htmlFor are form-scoped as \"{operationId}-%s\" to stay unique", name, len(ops), strings.Join(ops, ", "), name),
			Advice:  fmt.Sprintf("Do not target a bare id=%q in external selectors (E2E/CSS) — each form's input id is form-scoped (e.g. \"{operationId}-%s\"). Rename the field if a single shared id is intended.", name, name),
		})
	}
	return diags
}
