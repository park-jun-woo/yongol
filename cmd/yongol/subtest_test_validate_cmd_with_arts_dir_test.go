//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestValidateCmdWithArtsDir — WithArtsDir 서브테스트
package main

import (
	"bytes"
	"testing"
)

func subtestTestValidateCmdWithArtsDir(t *testing.T) {

	cmd := validateCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"../../examples/zenflow/opus4_7/specs", "/tmp/nonexistent-arts"})
	err := cmd.Execute()
	// Should still succeed — arts dir is optional and contract check skips on missing.
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}
