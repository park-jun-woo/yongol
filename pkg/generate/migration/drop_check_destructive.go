//ff:func feature=migration type=accessor control=sequence
//ff:what DropCheck.Destructive — CHECK 삭제는 비파괴적
package migration

func (op DropCheck) Destructive() bool { return false }
