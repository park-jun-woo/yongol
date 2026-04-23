//ff:func feature=migration type=util control=iteration dimension=1
//ff:what findLineCommentStart — 한 줄에서 `--` 시작 위치 반환 (single-quote 내부 무시)
package migration

// findLineCommentStart returns the byte index of the `--` that starts
// an SQL line comment, or -1 if the line has no such comment. Respects
// single-quoted string literals.
func findLineCommentStart(line string) int {
	scanner := lineCommentScanner{}
	for i := 0; i < len(line)-1; i++ {
		next, hit := scanner.step(line, i)
		if hit {
			return i
		}
		i = next
	}
	return -1
}
