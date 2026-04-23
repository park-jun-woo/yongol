//ff:func feature=migration type=accessor control=sequence
//ff:what CreateTable.Destructive — 항상 false (신규 생성)
package migration

func (op CreateTable) Destructive() bool { return false }
