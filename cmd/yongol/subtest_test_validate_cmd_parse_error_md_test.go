//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestValidateCmdParseErrorMD — ParseErrorMD 서브테스트
package main

import (
	"bytes"
	"testing"
)

func subtestTestValidateCmdParseErrorMD(t *testing.T) {

	tmpDir := t.TempDir()
	// Default (md) format drives the else branch (printParseErrors).
	if err := writeTestFile(tmpDir, "manifest.yaml", "::: not valid yaml :::\n\t- broken"); err != nil {
		t.Fatal(err)
	}
	cmd := validateCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{tmpDir})
	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected parse error, got nil")
	}
	if err.Error() != "parse failed" {
		t.Fatalf("expected 'parse failed', got: %v", err)
	}

}
