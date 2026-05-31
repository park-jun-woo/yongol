//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what applyRenames -- Op 내 FieldArg.Source / ResponseField.Source 참조를 rename 맵에 맞춰 갱신

package ir

// applyRenames updates FieldArg.Source and ResponseField.Source references
// in an Op to match renamed variables from previous shadowing resolution.
func applyRenames(op *Op, renames map[string]string) {
	if len(renames) == 0 {
		return
	}
	for _, args := range collectFieldArgSlices(op) {
		applyRenamesToFieldArgs(*args, renames)
	}
	// ResponseOp fields also reference variables via dot-notation.
	if op.Kind == OpResponse && op.Response != nil {
		applyRenamesToResponseFields(op.Response.Fields, renames)
	}
}
