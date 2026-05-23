//ff:type feature=agent type=helper
//ff:what validateLoopRoundResult — validate 루프 1회차 결과

package agent

// validateLoopRoundResult holds the outcome of one validate-fix round.
type validateLoopRoundResult struct {
	allErrors  int
	fixable    int
	roundFixed int
	done       bool // true when fixable == 0
}
