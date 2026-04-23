//ff:func feature=migration type=accessor control=sequence
//ff:what RenameColumn.Destructive — rename 은 비파괴적
package migration

func (op RenameColumn) Destructive() bool { return false }
