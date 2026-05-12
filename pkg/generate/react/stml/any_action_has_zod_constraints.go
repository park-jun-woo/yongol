//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 폼 액션 중 zod 제약조건이 있는 것이 있는지 확인한다
package stml

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// anyActionHasZodConstraints returns true if any form action has zod constraints.
func anyActionHasZodConstraints(actions []stmlparser.ActionBlock, constraints map[string]map[string]oapiparser.FieldConstraint) bool {
	if constraints == nil {
		return false
	}
	for _, a := range actions {
		if len(a.Fields) == 0 {
			continue
		}
		if fields := lookupConstraints(a.OperationID, constraints); len(fields) > 0 {
			return true
		}
	}
	return false
}
