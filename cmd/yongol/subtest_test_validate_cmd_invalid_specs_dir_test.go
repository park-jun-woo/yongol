//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestValidateCmdInvalidSpecsDir — InvalidSpecsDir 서브테스트
package main

import (
	"bytes"
	"testing"
)

func subtestTestValidateCmdInvalidSpecsDir(t *testing.T) {

	cmd := validateCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"/nonexistent/specs"})
	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error for nonexistent specs dir")
	}

}
