//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-54 — 단일 폼의 data-prefill 소스 실존(ERROR) 및 필드 커버리지(WARNING) 판정

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm54PrefillSourceForAction judges one action's data-prefill. The value must
// name a data-fetch operationId on the same page — the codegen reads that
// fetch's data variable (toLowerFirst(op)+"Data"), so a typo or a foreign-page
// operationId leaves the variable out of scope and the form cannot be wired
// (ERROR). For a resolvable source, each form requestBody field absent from the
// prefill 2xx response top-level (the responseFields judgment TM-20/TM-50 share)
// is a WARNING — the codegen fills it with a type-appropriate empty literal so
// the build passes, but that input opens blank.
func tm54PrefillSourceForAction(a stml.ActionBlock, file string, pageFetchOps map[string]bool, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if a.Prefill == "" {
		return nil
	}
	if !pageFetchOps[a.Prefill] {
		return []diagnostic.Diagnostic{{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-54] data-prefill %q on action %q is not a data-fetch operationId on this page — the prefill data source is out of scope", a.Prefill, a.OperationID),
			Advice:      fmt.Sprintf("Declare data-fetch=%q on the same page, or fix the data-prefill value to name an existing same-page fetch", a.Prefill),
			OperationID: a.OperationID,
		}}
	}
	entry, ok := opMap[a.Prefill]
	if !ok {
		return nil // TM-01 reports the unknown fetch operationId
	}
	resp := responseFields(entry.op)
	if len(resp) == 0 {
		return nil // untyped/void response — no field-coverage claim to make
	}
	return tm54FieldCoverage(a, file, resp)
}
