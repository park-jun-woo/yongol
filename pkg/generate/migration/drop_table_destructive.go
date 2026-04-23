//ff:func feature=migration type=accessor control=sequence
//ff:what DropTable.Destructive — 항상 true (데이터 손실)
package migration

func (op DropTable) Destructive() bool { return true }
