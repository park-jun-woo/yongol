//ff:func feature=migration type=test control=sequence
//ff:what TestSplitStateStepMethod — splitState.step 디스패치 커버
package migration

import "testing"

func TestSplitStateStepMethod(t *testing.T) {
	_ = splitStatements("CREATE TABLE \"X\" (a TEXT DEFAULT 'p;q' /* c */); CREATE INDEX i ON x(a);")
}
