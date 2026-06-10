//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-20 — data-capture 구문 위반 또는 응답 필드가 op의 OpenAPI 2xx 응답 스키마에 없음 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm20CaptureFieldInResponse checks an action's data-capture against the
// OpenAPI spec — the runtime twin of XOH-08 ([Captures] jsonpath must exist
// in the response). The raw attribute is re-parsed here (mirroring the
// ParseGuard / TM-17 split), so a syntax violation — including a sink
// outside auth.token / auth.refresh — is an ERROR. Each parsed respField
// must then be a top-level property of the operation's 2xx response
// schema; an unknown operationId is silently skipped (TM-02 reports it).
func tm20CaptureFieldInResponse(a stml.ActionBlock, file string, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if a.CaptureRaw == "" {
		return nil
	}
	binds, err := stml.ParseCapture(a.CaptureRaw)
	if err != nil {
		return []diagnostic.Diagnostic{{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-20] data-capture %q on action %q is invalid: %s", a.CaptureRaw, a.OperationID, err.Error()),
			Advice:      "Use data-capture=\"<respField> -> <sink>[, <respField> -> <sink>...]\" where sink is auth.token or auth.refresh",
			OperationID: a.OperationID,
		}}
	}
	entry, ok := opMap[a.OperationID]
	if !ok {
		return nil // TM-02 reports the unknown operationId
	}
	fields := responseFields(entry.op)
	var diags []diagnostic.Diagnostic
	for _, b := range binds {
		if _, ok := fields[b.RespField]; ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-20] data-capture field %q on action %q is not in the OpenAPI 2xx response schema of %q", b.RespField, a.OperationID, a.OperationID),
			Advice:      fmt.Sprintf("Add %q to the 2xx response schema of %q in the OpenAPI spec, or fix the data-capture field name", b.RespField, a.OperationID),
			OperationID: a.OperationID,
		})
	}
	return diags
}
