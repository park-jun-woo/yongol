//ff:func feature=migration type=accessor control=sequence
//ff:what AlterColumnDefault.Destructive — 기본값 변경은 비파괴적
package migration

func (op AlterColumnDefault) Destructive() bool { return false }
