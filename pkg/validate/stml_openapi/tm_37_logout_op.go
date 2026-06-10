//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-37 — data-logout operationId가 OpenAPI에 없거나 GET 메서드 (ERROR, TM-02/03 동형)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm37LogoutOp checks every layout data-logout that names a server-side
// session-ending operation (page-flow Phase010): the operationId must
// exist in OpenAPI and must not be a GET (the same contract TM-02/03
// enforce for data-action — ending a session is a mutation). A valueless
// data-logout declares no operation and is out of scope here (TM-38
// judges its mode fitness).
func tm37LogoutOp(layouts []stml.LayoutSpec, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, l := range layouts {
		if l.Logout == nil || l.Logout.OperationID == "" {
			continue
		}
		opID := l.Logout.OperationID
		entry, ok := opMap[opID]
		if !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:        l.File,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[TM-37] data-logout operationId %q in layout %q is not defined in OpenAPI", opID, l.Name),
				Advice:      "Name the server logout operation's OpenAPI operationId, or use a valueless data-logout for a client-only logout (bearer mode)",
				OperationID: opID,
			})
			continue
		}
		if entry.method == "GET" {
			diags = append(diags, diagnostic.Diagnostic{
				File:        l.File,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[TM-37] data-logout %q in layout %q references a GET endpoint; session-ending operations require POST/PUT/DELETE", opID, l.Name),
				Advice:      "Point data-logout at a state-changing operation (a GET must not end a session)",
				OperationID: opID,
			})
		}
	}
	return diags
}
