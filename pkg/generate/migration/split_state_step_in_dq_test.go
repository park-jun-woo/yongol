//ff:func feature=migration type=test control=sequence
//ff:what TestSplitStateStepInDQ — splitState.stepInDQ 이중인용 식별자 내부 커버
package migration

import "testing"

func TestSplitStateStepInDQMethod(t *testing.T) {
	_ = splitStatements("CREATE TABLE \"weird;name\" (id INT);")
}
