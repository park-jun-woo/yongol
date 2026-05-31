//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdParseError — ParseError 서브테스트
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func subtestTestGenerateCmdParseError(t *testing.T) {

	// A malformed manifest.yaml is detected then fails to parse, driving the
	// ParseDiagnostics branch (printParseErrors + "parse failed").
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("::: not valid yaml :::\n\t- broken"), 0o644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	_, _, err := runCmd(t, "generate", dir, t.TempDir())
	if err == nil {
		t.Fatal("expected parse error, got nil")
	}
	if !strings.Contains(err.Error(), "parse failed") {
		t.Fatalf("expected 'parse failed' error, got: %v", err)
	}

}
