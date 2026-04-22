//ff:type feature=gen-gogin type=model
//ff:what Limits — Q1 nesting depth + Q4 range body 순수 라인 상한

package qcheck

// Limits holds the Q1 (nesting depth) and Q4 (range body pure lines) ceilings.
// Generators that want to self-report depth/line excess pass these values.
type Limits struct {
	MaxDepth     int
	MaxPureLines int
}
