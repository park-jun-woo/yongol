//ff:func feature=migration type=util control=sequence
//ff:what appendPrimaryKeyClause — PRIMARY KEY (...) 절 추가 (비어 있으면 no-op)
package migration

import (
	"fmt"
	"strings"
)

// appendPrimaryKeyClause appends `,\n    PRIMARY KEY (col1, col2)` when
// pk is non-empty.
func appendPrimaryKeyClause(b *strings.Builder, pk []string) {
	if len(pk) == 0 {
		return
	}
	fmt.Fprintf(b, ",\n    PRIMARY KEY (%s)", strings.Join(pk, ", "))
}
