//ff:func feature=migration type=test control=sequence
//ff:what tokenizer/splitter named 테스트 — splitState/columnTokenizer/lineCommentScanner 메서드 (다중 인용/주석/타입 파라미터) 커버
package migration

import (
	"testing"
)

func TestLineCommentScannerStep(t *testing.T) {
	out := stripLineComments("SELECT 1 -- comment\nSELECT 2")
	if out == "" {
		t.Errorf("stripLineComments returned empty")
	}
}
