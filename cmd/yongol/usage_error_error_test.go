//ff:func feature=cli type=test control=sequence
//ff:what TestUsageErrorError — usageError.Error() 메시지 반환 검증

package main

import (
	"fmt"
	"testing"
)

func TestUsageErrorError(t *testing.T) {
	ue := &usageError{err: fmt.Errorf("bad args")}
	if got := ue.Error(); got != "bad args" {
		t.Errorf("Error() = %q, want %q", got, "bad args")
	}
}
