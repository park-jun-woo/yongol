//ff:func feature=validate type=rule control=sequence topic=policy-check
//ff:what checkOwnershipMapping — 한 @ownership 매핑에 대한 sqlc 쿼리 존재 검증

package query_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

// checkOwnershipMapping returns an XQP-30 diagnostic when the canonical
// OwnerLookup<Resource> query for om is missing from the have set. The
// second return signals whether a diagnostic was produced; false means
// the mapping was skipped (empty resource, duplicate, or already
// satisfied).
func checkOwnershipMapping(file string, om rego.OwnershipMapping, have, seen map[string]bool) (diagnostic.Diagnostic, bool) {
	if om.Resource == "" {
		return diagnostic.Diagnostic{}, false
	}
	want := ownerLookupName(om.Resource)
	key := file + "|" + want
	if seen[key] {
		return diagnostic.Diagnostic{}, false
	}
	seen[key] = true
	if have[want] {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:  file,
		Line:  om.SourceLine,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQP-30] @ownership %s — sqlc query %q not found; handler cannot load owner without it",
			om.Resource, want),
		Advice: buildAdvice(want, om.Table, om.Column, om.JoinTable, om.JoinFK),
	}, true
}
