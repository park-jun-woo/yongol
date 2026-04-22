//ff:func feature=validate type=util control=selection topic=ssac-structural
//ff:what crudType — seq.Type 이 CRUD (@get/@post/@put/@delete) 인지 판정

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// crudType reports whether the seq is one of @get/@post/@put/@delete.
func crudType(seq parsessac.Sequence) bool {
	switch seq.Type {
	case "get", "post", "put", "delete":
		return true
	}
	return false
}
