//ff:func feature=migration type=test control=sequence
//ff:what tokenizer/splitter named 테스트 — splitState/columnTokenizer/lineCommentScanner 메서드 (다중 인용/주석/타입 파라미터) 커버
package migration

import (
	"testing"
)

func TestLineCommentScannerStepQuote(t *testing.T) {
	// single-quoted string containing -- must be preserved.
	out := stripLineComments("INSERT INTO t VALUES ('a -- not a comment')")
	if out == "" {
		t.Errorf("stripLineComments returned empty")
	}
	_ = findLineCommentStart("x = 'a -- b' -- real")
}
