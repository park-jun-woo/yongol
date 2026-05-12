//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what enum 제약조건에서 z.enum() 체인을 생성한다
package stml

import (
	"fmt"
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// zodEnumChain builds a zod enum chain for an enum field constraint.
func zodEnumChain(fc oapiparser.FieldConstraint) string {
	quoted := make([]string, len(fc.Enum))
	for i, v := range fc.Enum {
		quoted[i] = fmt.Sprintf(`"%s"`, v)
	}
	base := fmt.Sprintf("z.enum([%s])", strings.Join(quoted, ", "))
	if !fc.Required {
		base += ".optional()"
	}
	return base
}
