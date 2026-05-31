//ff:func feature=rule type=test control=sequence
//ff:what TestCheckBaseSpecValidateExpectedError_ZeroCov — 헬퍼를 이름으로 직접 호출 (error + substring 만족 경로)

package rule

import (
	"errors"
	"testing"
)

func TestCheckBaseSpecValidateExpectedError_ZeroCov(t *testing.T) {
	// substring satisfied → no fatal.
	checkBaseSpecValidateExpectedError(t, errors.New("ruleID is required"), "ruleID")
	// empty wantSub → only nil-check matters.
	checkBaseSpecValidateExpectedError(t, errors.New("any error"), "")
}
