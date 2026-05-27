//ff:func feature=gen-fastapi type=util control=sequence
//ff:what emitSAImports — SQLAlchemy DML import 출력

package ssac

import (
	"fmt"
	"strings"
)

// emitSAImports writes the sqlalchemy DML imports.
func emitSAImports(b *strings.Builder, d importData) {
	var saImports []string
	if d.UsesSelect {
		saImports = append(saImports, "select")
	}
	if d.UsesUpdate {
		saImports = append(saImports, "update")
	}
	if d.UsesDelete {
		saImports = append(saImports, "delete")
	}
	if len(saImports) > 0 {
		b.WriteString(fmt.Sprintf("from sqlalchemy import %s\n", strings.Join(saImports, ", ")))
	}
}
