//ff:func feature=migration type=accessor control=sequence
//ff:what DropForeignKey.Destructive — FK 삭제는 비파괴적
package migration

func (op DropForeignKey) Destructive() bool { return false }
