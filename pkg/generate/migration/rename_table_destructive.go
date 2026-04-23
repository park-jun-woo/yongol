//ff:func feature=migration type=accessor control=sequence
//ff:what RenameTable.Destructive — rename 은 비파괴적
package migration

func (op RenameTable) Destructive() bool { return false }
