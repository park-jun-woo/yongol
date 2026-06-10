//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-33 보조 — 매핑 소스 검사 (무접두사 respField = 액션 op 2xx 응답 스키마, TM-20 인프라; route.* 면제)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm33CheckParamSource validates one redirect binding's source value. An
// unprefixed respField must be a top-level property of the action
// operation's 2xx response schema (responseFields — TM-20
// infrastructure); an unknown operationId is silently skipped (TM-02
// reports it). route.<Name> sources are exempt — they forward a
// current-page param, not response data. ParseRedirectParams already
// rejected item.* and empty sources.
func tm33CheckParamSource(p stml.LinkParamBind, a stml.ActionBlock, file string, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if strings.HasPrefix(p.Source, "route.") {
		return nil
	}
	entry, ok := opMap[a.OperationID]
	if !ok {
		return nil // TM-02 reports the unknown operationId
	}
	if _, ok := responseFields(entry.op)[p.Source]; ok {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:        file,
		Phase:       diagnostic.PhaseValidate,
		Level:       diagnostic.LevelError,
		Message:     fmt.Sprintf("[TM-33] data-redirect-params field %q on action %q is not in the OpenAPI 2xx response schema of %q — there is nothing to substitute at navigate time", p.Source, a.OperationID, a.OperationID),
		Advice:      fmt.Sprintf("Add %q to the 2xx response schema of %q in the OpenAPI spec, or fix the source field name", p.Source, a.OperationID),
		OperationID: a.OperationID,
	}}
}
