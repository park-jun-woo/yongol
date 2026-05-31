//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestValidateCmdInvalidFormat — InvalidFormat 서브테스트
package main

import (
	"bytes"
	"testing"
)

func subtestTestValidateCmdInvalidFormat(t *testing.T) {

	cmd := validateCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--format", "xml", "../../examples/zenflow/opus4_7/specs"})
	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error for invalid format")
	}

}
