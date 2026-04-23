//ff:func feature=migration type=accessor control=sequence
//ff:what AddCheck.SQL — ALTER TABLE ADD CONSTRAINT ... CHECK (...)
package migration

import "fmt"

func (op AddCheck) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s CHECK (%s);",
		op.Table, op.Check.Name, op.Check.Expression)
}
