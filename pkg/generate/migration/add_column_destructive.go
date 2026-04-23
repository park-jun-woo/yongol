//ff:func feature=migration type=accessor control=sequence
//ff:what AddColumn.Destructive — 항상 false (신규 컬럼)
package migration

func (op AddColumn) Destructive() bool { return false }
