//ff:func feature=migration type=test control=sequence
//ff:what TestColumnTokenizerStepInDQ — columnTokenizer.stepInDQ 이중인용 내부 커버
package migration

import "testing"

func TestColumnTokenizerStepInDQMethod(t *testing.T) {
	_ = tokenizeColumnDef(`"weird col" VARCHAR(10) NOT NULL`)
}
