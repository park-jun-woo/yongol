//ff:func feature=migration type=test control=sequence
//ff:what TestColumnTokenizerFlush — columnTokenizer.flush 버퍼 토큰 방출 커버
package migration

import "testing"

func TestColumnTokenizerFlushMethod(t *testing.T) {
	_ = tokenizeColumnDef(`note TEXT DEFAULT 'x' CHECK (note <> '')`)
}
