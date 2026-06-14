//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-57 — 상태 변경 mutation(POST/PUT/PATCH/DELETE) data-action이 data-redirect(성공 후 이동 대상)를 미선언 (ERROR, bearer 로그인 capture 액션 제외)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm57MutationRedirectRequired checks that every state-changing mutation
// declares where to navigate on success. "Where to go after create/update/
// delete" is an author decision, not a heuristic the codegen may invent
// (Phase040 decision #1): a mutation whose OpenAPI method is POST/PUT/PATCH/
// DELETE must carry a data-redirect, otherwise the generated CRUD screen
// stays on the same form after success and (for delete) refetches the
// deleted resource (BUG-132). A capture action (data-capture, the bearer
// login flow) is exempt — it commits the session and drives its own
// navigation, so the redirect is optional there. A GET data-action (a
// non-mutating action) and an unknown operationId (TM-02 reports it) yield
// no diagnostic. With a declared data-redirect the codegen's navigate
// branch always fires, guaranteeing the post-success move.
func tm57MutationRedirectRequired(a stml.ActionBlock, file string, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if a.CaptureRaw != "" {
		return nil // bearer login capture action — its own flow drives navigation
	}
	entry, ok := opMap[a.OperationID]
	if !ok {
		return nil // TM-02 reports the unknown operationId
	}
	if !isMutationMethod(entry.method) {
		return nil // GET data-action is non-mutating
	}
	if a.Redirect != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:        file,
		Phase:       diagnostic.PhaseValidate,
		Level:       diagnostic.LevelError,
		Message:     fmt.Sprintf("[TM-57] mutation action %q (%s) does not declare data-redirect (where to navigate on success)", a.OperationID, entry.method),
		Advice:      "Add data-redirect pointing to where the user should go on success: create→detail/list, update→detail, delete→list (a \"/\"-prefixed static path or an STML page-name reference)",
		OperationID: a.OperationID,
	}}
}
