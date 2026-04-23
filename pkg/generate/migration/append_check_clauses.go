//ff:func feature=migration type=util control=iteration dimension=1
//ff:what appendCheckClauses — CHECK 제약을 inline 절로 렌더
package migration

import (
	"fmt"
	"strings"
)

// appendCheckClauses writes `,\n    CONSTRAINT <name> CHECK (expr)` per
// check.
func appendCheckClauses(b *strings.Builder, checks []*CheckConstraint) {
	for _, chk := range checks {
		fmt.Fprintf(b, ",\n    CONSTRAINT %s CHECK (%s)", chk.Name, chk.Expression)
	}
}
