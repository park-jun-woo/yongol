//ff:func feature=manifest type=test control=sequence
//ff:what extractDefaultString — DEFAULT 'draft' 문자열 리터럴 추출

package ddl

import "testing"

func TestExtractDefaultString_Match(t *testing.T) {
	if got := extractDefaultString("status VARCHAR(32) NOT NULL DEFAULT 'draft'"); got != "draft" {
		t.Errorf("got %q, want draft", got)
	}
}
