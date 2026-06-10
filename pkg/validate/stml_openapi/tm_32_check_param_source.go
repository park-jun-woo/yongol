//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-32 보조 — 매핑 소스 검사 (item.* each 컨텍스트/스키마 = TM-30 인프라, route.* 자기 라우트 = TM-27 인프라)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm32CheckParamSource validates one binding's source value. item.<Field>
// is only legal inside a data-each whose item schema has the field (TM-30
// infrastructure; a nil schema stays silent — TM-01/TM-07 territory).
// route.<Name> must be a segment of this page's own resolved route (TM-27
// infrastructure, case-exact). ParseLinkParams already rejected any other
// prefix.
func tm32CheckParamSource(p stml.LinkParamBind, ref linkRefCtx, file string, ownRouteParams map[string]bool) []diagnostic.Diagnostic {
	field, isItem := itemParamField(p.Source)
	if isItem && !ref.InEach {
		return []diagnostic.Diagnostic{{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-32] data-link param source %q is only valid inside a data-each block — no row item is in scope here", p.Source),
			Advice:  "Move the link inside the data-each block whose rows it targets, or use a route.<Name> source",
		}}
	}
	if isItem && ref.ItemFields != nil && !ref.ItemFields[field] {
		return []diagnostic.Diagnostic{{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-32] data-link param source %q references field %q which is not in the item schema of the enclosing data-each array", p.Source, field),
			Advice:  fmt.Sprintf("Add %q to the array item schema in the OpenAPI response, or fix the item.<Field> source", field),
		}}
	}
	if isItem {
		return nil
	}
	name := strings.TrimPrefix(p.Source, "route.")
	if name != p.Source && !ownRouteParams[name] {
		return []diagnostic.Diagnostic{{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-32] data-link param source %q is not a segment of this page's resolved route — the value is always undefined at runtime", p.Source),
			Advice:  fmt.Sprintf("Add `:%s` to this page's route (data-route or route.* consumption) or fix the source name", name),
		}}
	}
	return nil
}
