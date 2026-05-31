//ff:func feature=migration type=test control=sequence
//ff:what TestSplitStateStepInSQ — splitState.stepInSQ 단일인용 문자열 내부 커버
package migration

import "testing"

func TestSplitStateStepInSQMethod(t *testing.T) {
	_ = splitStatements("INSERT INTO t VALUES ('a; b ''q'' c'); SELECT 1;")
}
