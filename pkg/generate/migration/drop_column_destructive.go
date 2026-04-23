//ff:func feature=migration type=accessor control=sequence
//ff:what DropColumn.Destructive — 항상 true (데이터 손실)
package migration

func (op DropColumn) Destructive() bool { return true }
