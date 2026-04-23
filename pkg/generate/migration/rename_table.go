//ff:type feature=migration type=model
//ff:what RenameTable — ALTER TABLE RENAME TO Operation (@rename 힌트 소스)
package migration

type RenameTable struct {
	From, To string
}
