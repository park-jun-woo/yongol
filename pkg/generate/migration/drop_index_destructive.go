//ff:func feature=migration type=accessor control=sequence
//ff:what DropIndex.Destructive — 인덱스 삭제는 비파괴적 (데이터 손실 없음)
package migration

func (op DropIndex) Destructive() bool { return false }
