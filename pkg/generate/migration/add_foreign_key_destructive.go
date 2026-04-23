//ff:func feature=migration type=accessor control=sequence
//ff:what AddForeignKey.Destructive — FK 추가는 비파괴적
package migration

func (op AddForeignKey) Destructive() bool { return false }
