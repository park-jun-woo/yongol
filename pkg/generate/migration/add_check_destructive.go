//ff:func feature=migration type=accessor control=sequence
//ff:what AddCheck.Destructive — CHECK 추가는 비파괴적
package migration

func (op AddCheck) Destructive() bool { return false }
