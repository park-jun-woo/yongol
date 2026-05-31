//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestValidateCmdJSONFormat — JSONFormat 서브테스트
package main

import (
	"bytes"
	"testing"
)

func subtestTestValidateCmdJSONFormat(t *testing.T) {

	cmd := validateCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--format", "json", "../../examples/zenflow/opus4_7/specs"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}
