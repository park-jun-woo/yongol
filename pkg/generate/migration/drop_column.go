//ff:type feature=migration type=model
//ff:what DropColumn — ALTER TABLE DROP COLUMN Operation
package migration

type DropColumn struct {
	Table, Column    string
	AllowDestructive bool
}
