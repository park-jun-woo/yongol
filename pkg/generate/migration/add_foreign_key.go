//ff:type feature=migration type=model
//ff:what AddForeignKey — ALTER TABLE ADD CONSTRAINT FOREIGN KEY Operation
package migration

type AddForeignKey struct {
	Table string
	FK    *ForeignKey
}
