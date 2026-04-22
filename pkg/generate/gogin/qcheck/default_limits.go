//ff:func feature=gen-gogin type=util control=sequence
//ff:what DefaultLimits — filefunc 기본 Q1/Q4 budget 반환 (depth ≤ 3, pure ≤ 10)

package qcheck

// DefaultLimits returns the filefunc baseline budget: depth ≤ 3 (iteration
// dimension=2), range body pure lines ≤ 10. Callers override when the
// template advertises a wider dimension (dimension+1 depth budget).
func DefaultLimits() Limits {
	return Limits{MaxDepth: 3, MaxPureLines: 10}
}
