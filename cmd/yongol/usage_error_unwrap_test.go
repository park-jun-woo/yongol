//ff:func feature=cli type=test control=sequence
//ff:what TestUsageErrorUnwrap — usageError.Unwrap() 원본 에러 반환 검증

package main

import (
	"fmt"
	"testing"
)

func TestUsageErrorUnwrap(t *testing.T) {
	inner := fmt.Errorf("inner error")
	ue := &usageError{err: inner}
	if got := ue.Unwrap(); got != inner {
		t.Errorf("Unwrap() = %v, want %v", got, inner)
	}
}
