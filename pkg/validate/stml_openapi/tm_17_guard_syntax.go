//ff:func feature=validate type=rule control=sequence dimension=1 topic=stml-openapi
//ff:what TM-17 — data-state 가드 문자열이 §3.4 EBNF에 적합한지 검증 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm17GuardSyntax checks a single data-state condition that uses guard
// combinators (&&, ||, leading !, or parentheses) against the §3.4 EBNF. A
// parse failure yields a TM-17 ERROR. Legacy non-combinator conditions
// (field=value, .loading/.error/.empty, bare field) are skipped and return nil.
func tm17GuardSyntax(condition, file string) []diagnostic.Diagnostic {
	if !stml.HasGuardCombinator(condition) {
		return nil
	}
	if _, err := stml.ParseGuard(condition); err != nil {
		return []diagnostic.Diagnostic{{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-17] data-state guard %q is not valid guard syntax: %s", condition, err.Error()),
			Advice:  "Use only comparisons (model.Field = value), lifecycle (model.Field.loading/error/empty), logical && / ||, negation !, and parentheses. Function calls, arithmetic, and ternaries are not allowed.",
		}}
	}
	return nil
}
