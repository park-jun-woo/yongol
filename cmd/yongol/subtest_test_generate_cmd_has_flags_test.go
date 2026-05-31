//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdHasFlags — HasFlags 서브테스트
package main

import (
	"testing"
)

func subtestTestGenerateCmdHasFlags(t *testing.T) {

	cmd := generateCmd()
	if cmd.Flags().Lookup("backend") == nil {
		t.Error("expected --backend flag")
	}
	if cmd.Flags().Lookup("frontend") == nil {
		t.Error("expected --frontend flag")
	}

}
