//ff:func feature=migration type=accessor control=sequence
//ff:what AlterColumnType.Destructive — 타입 변경은 항상 파괴적
package migration

func (op AlterColumnType) Destructive() bool { return true }
