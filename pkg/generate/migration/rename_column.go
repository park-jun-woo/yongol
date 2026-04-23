//ff:type feature=migration type=model
//ff:what RenameColumn — ALTER TABLE RENAME COLUMN Operation (@rename 힌트 소스)
package migration

type RenameColumn struct {
	Table, From, To string
}
