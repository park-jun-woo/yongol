//ff:func feature=rule type=test-helper control=sequence
//ff:what checkBaseSpecValidateExpectedError — BaseSpec.Validate 가 error + wantSub 기대치를 만족하는지 확인

package rule

import (
	"strings"
	"testing"
)

func checkBaseSpecValidateExpectedError(t *testing.T, err error, wantSub string) {
	t.Helper()
	if err == nil {
		t.Fatalf("Validate() = nil; want error containing %q", wantSub)
	}
	if wantSub != "" && !strings.Contains(err.Error(), wantSub) {
		t.Fatalf("Validate() error = %q; want substring %q", err.Error(), wantSub)
	}
}
