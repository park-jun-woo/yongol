//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestValidateCmdValidSpecs — ValidSpecs 서브테스트
package main

import (
	"bytes"
	"testing"
)

func subtestTestValidateCmdValidSpecs(t *testing.T) {

	cmd := validateCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"../../examples/zenflow/opus4_7/specs"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}
