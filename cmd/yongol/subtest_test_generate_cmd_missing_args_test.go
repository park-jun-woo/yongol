//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdMissingArgs — MissingArgs 서브테스트
package main

import (
	"testing"
)

func subtestTestGenerateCmdMissingArgs(t *testing.T) {

	_, _, err := runCmd(t, "generate")
	if err == nil {
		t.Fatal("expected usage error for missing args, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}

}
