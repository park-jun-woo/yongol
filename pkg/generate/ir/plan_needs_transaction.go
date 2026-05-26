//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what planNeedsTransaction -- Op 목록에 mutating 연산 존재 여부 판정

package ir

// planNeedsTransaction returns true when at least one mutating op
// (@post/@put/@delete) is present.
func planNeedsTransaction(ops []Op) bool {
	for _, op := range ops {
		switch op.Kind {
		case OpPost, OpPut, OpDelete:
			return true
		}
	}
	return false
}
