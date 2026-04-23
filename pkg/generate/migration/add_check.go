//ff:type feature=migration type=model
//ff:what AddCheck — ALTER TABLE ADD CONSTRAINT CHECK Operation
package migration

type AddCheck struct {
	Table string
	Check *CheckConstraint
}
