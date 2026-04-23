//ff:type feature=migration type=model
//ff:what DropCheck — ALTER TABLE DROP CONSTRAINT <check> Operation
package migration

type DropCheck struct{ Table, Name string }
