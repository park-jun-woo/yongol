//ff:func feature=migration type=test control=sequence
//ff:what TestLineCommentScannerStep — lineCommentScanner.step 라인 코멘트 스캔 커버
package migration

import "testing"

func TestLineCommentScannerStepMethod(t *testing.T) {
	if findLineCommentStart("SELECT 1 -- c") < 0 {
		t.Errorf("expected to find comment start")
	}
}
