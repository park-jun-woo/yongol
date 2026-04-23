//ff:func feature=migration type=accessor control=sequence
//ff:what CreateIndex.Destructive — 인덱스 생성은 비파괴적
package migration

func (op CreateIndex) Destructive() bool { return false }
