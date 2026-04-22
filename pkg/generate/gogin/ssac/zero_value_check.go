//ff:func feature=gen-gogin type=util control=sequence
//ff:what zeroValueCheck — 타입에 따른 zero-value 비교 표현식

package ssac

// zeroValueCheck returns the Go zero-value comparison for @empty/@exists.
// int64 → "== 0", string → `== ""`, default → "== 0" (ID 기준).
func zeroValueCheck(target string) string {
	return target + " == 0"
}


