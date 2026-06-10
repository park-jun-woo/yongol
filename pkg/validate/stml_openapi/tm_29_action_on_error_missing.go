//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-29 — 에러 응답(4xx/5xx)을 선언한 op 를 소비하는 data-action 블록에 data-on-error 부재 (WARNING, 기본 인라인 슬롯 폴백)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm29ActionOnErrorMissing warns when a data-action block consumes an
// operation that declares at least one 4xx/5xx response but the block has
// no data-on-error element (BUG-113 (5) — error-UX omission made visible
// at the SSOT level). The page-flow Phase004 default slot (role="alert"
// next to the submit button) guarantees a baseline display, so the level
// is WARNING, not ERROR: the decision of *where* the server error appears
// is what the author is asked to make explicit. Response body schema shape
// is irrelevant — only the declared 4xx/5xx keys matter (the Phase003
// message-extraction fallback absorbs non-ErrorResponse bodies). An
// unknown operationId is silently skipped (TM-02 reports it); data-fetch
// (GET) blocks are out of scope by construction (per-action rule).
func tm29ActionOnErrorMissing(a stml.ActionBlock, file string, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if a.OnErrorNode {
		return nil
	}
	entry, ok := opMap[a.OperationID]
	if !ok {
		return nil // TM-02 reports the unknown operationId
	}
	if !opDeclaresErrorResponse(entry.op) {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:        file,
		Phase:       diagnostic.PhaseValidate,
		Level:       diagnostic.LevelWarning,
		Message:     fmt.Sprintf("[TM-29] operation %q declares error responses but the action block has no `data-on-error`; a default inline message is shown — add `data-on-error` to choose where the server error appears", a.OperationID),
		Advice:      "Add a `data-on-error` element inside the action block to decide where the server error message is displayed",
		OperationID: a.OperationID,
	}}
}
