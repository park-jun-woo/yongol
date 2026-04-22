//ff:func feature=gen-gogin type=util control=sequence
//ff:what clampDepth — 음수 depth를 0으로 보정 (잘못 매칭된 }로 인한 underflow 방지)

package ffannot

// clampDepth returns d when non-negative, else 0. Used after summing a brace
// delta so a spurious unmatched '}' (e.g. inside a raw string literal we don't
// parse) can't push depth below the function body baseline.
func clampDepth(d int) int {
	if d < 0 {
		return 0
	}
	return d
}
