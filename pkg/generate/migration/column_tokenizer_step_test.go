//ff:func feature=migration type=test control=sequence
//ff:what TestColumnTokenizerStep — columnTokenizer.step 디스패치 커버
package migration

import "testing"

func TestColumnTokenizerStepMethod(t *testing.T) {
	_ = tokenizeColumnDef(`status VARCHAR(20) NOT NULL DEFAULT 'a''b'`)
}
