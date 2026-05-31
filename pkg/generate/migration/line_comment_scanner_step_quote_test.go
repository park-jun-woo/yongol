//ff:func feature=migration type=test control=sequence
//ff:what TestLineCommentScannerStepQuote — lineCommentScanner.stepQuote 인용 내부 -- 보존 커버
package migration

import "testing"

func TestLineCommentScannerStepQuoteMethod(t *testing.T) {
	// -- inside a single-quoted string is not a comment.
	if findLineCommentStart("x = 'a -- b'") >= 0 {
		t.Errorf("-- inside quotes should not be a comment start")
	}
	_ = findLineCommentStart("x = 'a -- b' -- real")
}
