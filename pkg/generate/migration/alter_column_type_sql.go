//ff:func feature=migration type=accessor control=selection
//ff:what AlterColumnType.SQL — ALTER TABLE ALTER COLUMN TYPE USING
package migration

import (
	"fmt"
	"strings"
)

func (op AlterColumnType) SQL() string {
	if op.Using != "" {
		return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s USING %s;",
			op.Table, op.Column, op.To.SQL(), op.Using)
	}
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s USING %s::%s;",
		op.Table, op.Column, op.To.SQL(), op.Column, strings.ToLower(op.To.Base))
}
