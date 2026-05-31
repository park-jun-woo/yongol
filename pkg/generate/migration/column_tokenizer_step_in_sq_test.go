//ff:func feature=migration type=test control=sequence
//ff:what TestColumnTokenizerStepInSQ — columnTokenizer.stepInSQ 단일인용 내부 커버
package migration

import "testing"

func TestColumnTokenizerStepInSQMethod(t *testing.T) {
	_ = tokenizeColumnDef(`status TEXT DEFAULT 'a; b ''q'' c'`)
}
