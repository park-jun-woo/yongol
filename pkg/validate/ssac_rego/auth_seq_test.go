//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func authSeq(action, resource string, line int) ssac.Sequence {
	return ssac.Sequence{Type: "auth", Action: action, Resource: resource, Line: line}
}
