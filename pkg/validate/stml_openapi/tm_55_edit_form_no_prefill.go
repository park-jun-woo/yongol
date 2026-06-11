//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-55 — GET-by-id fetch가 있는데 PUT/PATCH 폼이 data-prefill 미선언 (WARNING, 빈 폼 안내)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm55EditFormNoPrefill warns on the canonical edit page that forgets to wire
// prefill (plans/gen/frontend Phase035, BUG-124): the page has a GET-by-id
// data-fetch (reads the current values) and a PUT/PATCH data-action carrying
// data-field inputs, yet declares no data-prefill — so the form is generated
// empty and a single-field edit forces re-entering every field. It makes that
// blank-form generation visible and points to data-prefill. Forms that already
// declare data-prefill, field-less actions, and non-PUT/PATCH actions are
// silent; an unknown operationId is TM-02's finding.
func tm55EditFormNoPrefill(page stml.PageSpec, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if !pageHasGetByIdFetch(page, opMap) {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, a := range stml.CollectChildActions(page.Children) {
		if a.Prefill != "" || len(a.Fields) == 0 {
			continue
		}
		entry, ok := opMap[a.OperationID]
		if !ok || (entry.method != "PUT" && entry.method != "PATCH") {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        page.FileName,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelWarning,
			Message:     fmt.Sprintf("[TM-55] %s edit form %q has no data-prefill though the page has a GET-by-id fetch — the form is generated empty", entry.method, a.OperationID),
			Advice:      "Add data-prefill=\"<the GET-by-id fetch operationId>\" to the form so it opens with the current values",
			OperationID: a.OperationID,
		})
	}
	return diags
}
