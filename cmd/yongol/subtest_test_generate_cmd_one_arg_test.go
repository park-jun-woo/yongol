//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdOneArg — OneArg 서브테스트
package main

import (
	"testing"
)

func subtestTestGenerateCmdOneArg(t *testing.T) {

	_, _, err := runCmd(t, "generate", "/tmp/specs-only")
	if err == nil {
		t.Fatal("expected usage error for single arg, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}

}
