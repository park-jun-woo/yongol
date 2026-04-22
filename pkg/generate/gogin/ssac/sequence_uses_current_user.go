//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sequenceUsesCurrentUser — 단일 Sequence에 currentUser 참조가 있는지 검사

package ssac

import (
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// sequenceUsesCurrentUser returns true when seq is an @auth step (which
// always consumes currentUser.ID / currentUser.Role) or when any Input
// value is prefixed with "currentUser." — extracted so needsCurrentUser
// keeps a flat loop at depth 1.
func sequenceUsesCurrentUser(seq ssacparser.Sequence) bool {
	// @auth always uses currentUser.ID / currentUser.Role
	if seq.Type == "auth" {
		return true
	}
	for _, v := range seq.Inputs {
		if strings.HasPrefix(v, "currentUser.") {
			return true
		}
	}
	return false
}
