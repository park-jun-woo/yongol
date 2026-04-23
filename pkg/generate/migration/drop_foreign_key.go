//ff:type feature=migration type=model
//ff:what DropForeignKey — ALTER TABLE DROP CONSTRAINT <fk> Operation
package migration

type DropForeignKey struct{ Table, Name string }
