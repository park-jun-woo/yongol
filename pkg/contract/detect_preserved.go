//ff:func feature=contract type=util control=sequence
//ff:what DetectPreserved — 파일의 //ff:checked hash 와 재계산 hash 를 비교해 preserve 상태 판정

package contract

import (
	"os"
	"strings"
)

// DetectPreserved reads filePath, locates the first `//ff:checked`
// annotation, and returns a PreservedState by comparing the saved
// hash with ComputeBodyHash applied to the same bytes.
//
// Returns StateNotApplicable when no `//ff:checked` annotation is
// present — callers treat these as out-of-scope for preserve drift.
func DetectPreserved(filePath string) (PreservedState, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return StateNotApplicable, err
	}
	m := checkedLineRe.FindSubmatch(src)
	if m == nil {
		return StateNotApplicable, nil
	}
	saved := strings.ToLower(string(m[1]))
	current := ComputeBodyHash(src)
	if current == "" {
		return StateNotApplicable, nil
	}
	if saved == strings.ToLower(current) {
		return StateUntouched, nil
	}
	return StatePreserved, nil
}
