//ff:func feature=migration type=test control=sequence
//ff:what TestColumnTokenizerFinish — columnTokenizer.finish 마지막 토큰 처리 커버
package migration

import "testing"

func TestColumnTokenizerFinishMethod(t *testing.T) {
	toks := tokenizeColumnDef(`id BIGINT PRIMARY KEY`)
	if len(toks) == 0 {
		t.Errorf("expected tokens")
	}
}
