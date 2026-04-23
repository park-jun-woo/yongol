//ff:func feature=tsx-parser type=accessor control=iteration dimension=1
//ff:what (v *visitor).resolve — swc 1-based byte offset → 1-based (line, col)

package tsx

// resolve converts a swc span.start (1-based byte offset) into 1-based
// (line, col). A start of 0 or out-of-range yields (0, 0) so callers can
// detect missing positions.
func (v *visitor) resolve(spanStart int) (line, col int) {
	if spanStart <= 0 {
		return 0, 0
	}
	// swc spans are 1-based; convert to 0-based byte offset.
	off := spanStart - 1
	if off > len(v.src) {
		off = len(v.src)
	}
	// Binary search over lineOffset.
	lo, hi := 0, len(v.lineOffset)-1
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if v.lineOffset[mid] <= off {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	line = lo + 1
	col = off - v.lineOffset[lo] + 1
	return
}
