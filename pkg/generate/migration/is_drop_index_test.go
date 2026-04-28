//ff:func feature=migration type=test-helper control=sequence
//ff:what isDropIndex — op 이 DropIndex 타입인지 판별

package migration

// isDropIndex reports whether op is a DropIndex operation. Isolated as a
// helper so TestDiff_IndexMethodChange_BtreeToGIN avoids an if-at-depth-2
// inside its for loop (Q1).
func isDropIndex(op Operation) bool {
	_, ok := op.(DropIndex)
	return ok
}
