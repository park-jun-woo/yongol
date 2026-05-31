//ff:func feature=migration type=test control=sequence
//ff:what TestSplitStateStepInBC — splitState.stepInBC 블록 코멘트 내부 커버
package migration

import "testing"

func TestSplitStateStepInBCMethod(t *testing.T) {
	_ = splitStatements("SELECT 1 /* block ; comment */ ; SELECT 2;")
}
